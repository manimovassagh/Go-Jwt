package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jwt-project/controllers"
	"github.com/jwt-project/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
func main() {

	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
