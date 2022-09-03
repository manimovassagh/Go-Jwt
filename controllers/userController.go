package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/jwt-project/initializers"
	"github.com/jwt-project/models"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	// get The Email and password of request body
	var body struct {
		Email    string `gorm:"unique"`
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "OOpss You have Error in parsing Body",
		})

		return
	}
	//Hash the  password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to Hash Password",
		})
		return
	}
	// Create The user
	user := models.User{Email: body.Email, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":          "Failed to create user",
			"reson of error": "Duplicate Email is not acceptable",
		})
		return
	}
	//Respond

	c.JSON(http.StatusOK, gin.H{
		"Success": "User Created ",
	})

}

func Login(c *gin.Context) {

	// get The Email and password of request body
	var body struct {
		Email    string `gorm:"unique"`
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "OOpss You have Error in parsing Body",
		})

		return
	}

	//Lookup the user base on email and password
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return

	}

	//compare set in  pass with saved pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Email or Password",
		})
		return
	}
	//gernerate a jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"nbf": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		fmt.Println(err)
		return
	}

	//send it back as normal token instead of cookie
	//c.JSON(http.StatusOK, gin.H{
	//	"token": tokenString,
	//})

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": "Cookie Created",
	})
}

func Validate(c *gin.Context) {
c.JSON(http.StatusOK,gin.H{
	"message":"This user is Logged In",
})
}
