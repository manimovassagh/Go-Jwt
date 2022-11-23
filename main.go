package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jwt-project/controllers"
	"github.com/jwt-project/initializers"
	"github.com/jwt-project/middleware"
)

func init() {
	initializers.ConnectToDb()
	initializers.LoadEnvVariables()
	initializers.SyncDatabase()
}
func main() {

	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/Just Test", controllers.Validate)
	r.Run()

}
