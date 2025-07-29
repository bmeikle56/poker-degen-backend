package main

import (
	"os"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
	"pokerdegen/handlers"
	"pokerdegen/database"
	"pokerdegen/middleware"
	"log"
	_ "github.com/lib/pq"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	godotenv.Load()

	err := database.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	r := gin.Default()

	r.POST("/login", handlers.LoginHandler)
	r.POST("/signup", handlers.SignupHandler)
	r.POST("/deleteAccount", handlers.DeleteAccountHandler)
	r.POST("/modelWrapper", middleware.AuthMiddleware(), handlers.ModelWrapperHandler)
	r.Run(":" + port)
}
