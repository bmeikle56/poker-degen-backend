package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"pokerdegen/services"
	"pokerdegen/models"
)

func LoginHandler(c *gin.Context) {
	var req models.AuthRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	
	err := services.LoginService(req.Username, req.Password)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}