package main

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"poker-degen/handlers"
	"poker-degen/middleware"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	r := gin.Default()

	r.POST("/login", middleware.AuthMiddleware(), handlers.LoginHandler)
	r.Run(":" + port)
}
