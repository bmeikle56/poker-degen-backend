package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"poker-degen/services"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(c *gin.Context) {
	var req loginRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	
	success := services.LoginService(req.Username, req.Password)

	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Login failed"})
	}
}