package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp() {

	// get The Email and password of request body
	var body struct {
		Email    string
		Password string
	}

	c.Bind(&body) != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"OOpss You have Error in parsin Body"
		}),
	}
	//Hash the  password
	hash,err:=bcrypt.GenerateFormPassword([]byte(body.Password),10)

	if err !=nil{
		fmt.println("You have error in Hashing password")
	}
	// Create The user

	//Respond

}