package main

import (
	"fmt"
	"gin-jwt/controllers"
	"gin-jwt/initializers"
	"gin-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	fmt.Println("hello")
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.Use(middleware.RequireAuth)
	r.GET("/validate", controllers.Validate)
	r.Run()
}
