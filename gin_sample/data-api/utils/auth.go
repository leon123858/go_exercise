package utils

import (
	"encoding/json"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
)

type authHeader struct {
	IDToken string `header:"Authorization"`
}

func AuthUser(c *gin.Context) {
	header := authHeader{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
			"text":  "Must provide Authorization header with format `token`",
		})
		c.Abort()
		return
	}

	token, err := AuthClient.VerifyIDToken(AuthCtx, header.IDToken)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"error": err.Error(),
			"text":  "error verifying ID token",
		})
		c.Abort()
		return
	}
	jsonData, _ := json.Marshal(token)
	c.Set("token", string(jsonData))
}

func GetUserToken(c *gin.Context) auth.Token {
	value := []byte(c.GetString("token"))
	var token auth.Token
	json.Unmarshal(value, &token)
	return token
}
