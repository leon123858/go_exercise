package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/leon123858/go_Exercise/session/middleware"
)

func main() {
	e := echo.New()

	// e.Use(middleware.SessionMiddleware())

	e.GET("/", func(c echo.Context) error {
		// return string include information of session
		return c.String(http.StatusOK, fmt.Sprintf("Hello %v", c.Get("session").(*middleware.UserSession).UserName))
	}, middleware.SessionMiddleware())

	e.GET("/login", func(c echo.Context) error {
		// get user info from database
		// skip
		// set session
		userSession := &middleware.UserSession{
			UserId:   1,
			UserName: "Leon",
		}
		if err := middleware.AddSession(c, userSession); err != nil {
			c.Logger().Error(err)
			return err
		}
		return c.JSON(http.StatusOK, "login success")
	})

	e.GET("/logout", func(c echo.Context) error {
		// delete session
		if err := middleware.DeleteSession(c); err != nil {
			c.Logger().Error(err)
			return err
		}
		return c.JSON(http.StatusOK, "logout success")
	}, middleware.SessionMiddleware())

	e.Start(":8080")
}
