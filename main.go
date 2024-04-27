package main

import (
	"fmt"

	"github.com/Shenon69/jwt-go/controllers"
	"github.com/Shenon69/jwt-go/initializers"
	"github.com/Shenon69/jwt-go/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	fmt.Println("Initializing the application...")
	initializers.LoadEnv()
	initializers.ConnectDB()
	initializers.SyncDatabase()
	fmt.Println("Application initialized🚀")
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}