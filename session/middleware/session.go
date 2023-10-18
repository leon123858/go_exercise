package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/leon123858/go_Exercise/session/utils"
)

type UserSession struct {
	UserId   int
	UserName string
}

// create a middleware for echo
func SessionMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get sessionId from cookie
			sessionId, err := c.Cookie("sessionId")
			if err != nil {
				c.Logger().Error(err)
				c.JSON(403, "no session id")
				return err
			}

			userSession := &UserSession{}

			// get session from redis
			if err := utils.MyCache.Get(context.TODO(), sessionId.Value, &userSession); err != nil {
				c.Logger().Error(err)
				c.JSON(403, "no session")
				return err
			}

			// set session to echo context
			c.Set("session", userSession)
			c.Set("sessionId", sessionId.Value)

			// call next handler
			return next(c)
		}
	}
}

// add session into user cookie
func AddSession(c echo.Context, userSession *UserSession) error {
	// create a new session id
	sessionId := uuid.New().String()

	// expire time is 1 hour
	expire := time.Hour

	// set session to redis
	if err := utils.MyCache.Set(&cache.Item{
		Key:   sessionId,
		Value: userSession,
		TTL:   expire,
	}); err != nil {
		return err
	}

	// set session id to cookie
	cookie := &http.Cookie{
		Name:     "sessionId",
		Value:    sessionId,
		Expires:  time.Now().Add(expire),
		HttpOnly: true,
	}

	c.SetCookie(cookie)
	return nil
}

// delete session from user cookie

func DeleteSession(c echo.Context) error {
	// get sessionId from context session
	sessionId := c.Get("sessionId").(string)

	// delete session from redis
	if err := utils.MyCache.Delete(context.TODO(), sessionId); err != nil {
		return err
	}

	// remove session id from cookie
	cookie := &http.Cookie{
		Name:     "sessionId",
		Value:    "",
		Expires:  time.Now().Add(-1),
		HttpOnly: true,
	}

	c.SetCookie(cookie)
	return nil
}
