package main

import (
    "net/http"
		"os"
		"github.com/joho/godotenv"
		"github.com/gin-gonic/gin"
)

func main() {
    godotenv.Load()

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    
		r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello, Gin!")
    })

    r.POST("/chat", func(c *gin.Context) {
        var json struct {
            Message string `json:"message"`
        }

        if err := c.BindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
            return
        }

        // Just echo back the message for now
        c.JSON(http.StatusOK, gin.H{
            "reply": "You said: " + json.Message,
        })
    })

    r.Run()
}
