package handlers

import (
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"pokerdegen/services"
	"pokerdegen/models"
)

func SignupHandler(c *gin.Context) {
	var req models.AuthRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	
	err := services.SignupService(req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to signup user",
			"details": err.Error(),
		})
	} else {
		authToken := os.Getenv("AUTH_TOKEN")
		c.JSON(http.StatusOK, gin.H{
			"token": authToken,
			"message": "Sign up successful",
		})
	}
}