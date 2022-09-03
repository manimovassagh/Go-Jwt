package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("In middleware")

	//Get The cookie from request
	tokenString, err := c.Cookie("Autorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)

	}

	//Decode/validate cookie
	//parse takes the
	//check the expiration

	//find the user with token sub

	//continue

	c.Next()
}
