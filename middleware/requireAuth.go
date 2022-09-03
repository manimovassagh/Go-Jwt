package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jwt-project/initializers"
	"github.com/jwt-project/models"
)

func RequireAuth(c *gin.Context) {
	fmt.Println("In middleware")

	//Get The cookie from request
	tokenString, err := c.Cookie("Autorization")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)

	}

	//Decode/validate cookie
	// Parse takes the token string and a function for looking up the key. The latter is especially
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		//check the expiration
		if float64(time.Now().Unix()) >["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//find the user with token sub
		var user models.User
		initializers.DB.First(&user,claims["sub"])
		if user.ID==0{

		}
		//Attach to requests
		//continue

		c.Next()
		fmt.Println(claims["foo"], claims["nbf"])
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	
}
