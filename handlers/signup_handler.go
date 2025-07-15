package handlers

import (
	"net/http"
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
		c.JSON(http.StatusOK, gin.H{"message": "Sign up successful"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}