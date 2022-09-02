package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
			"error": "OOpss You have Error in parsin Body",
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
			"error": "Failed to create user",
		})
		return
	}
	//Respond

	c.JSON(http.StatusOK, gin.H{
		"Success": "User Created ",
	})

}
