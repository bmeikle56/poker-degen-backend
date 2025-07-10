package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"poker-degen/services"
	"poker-degen/models"
)

func ModelWrapperHandler(c *gin.Context) {
	var req models.ModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input", "details": err.Error()})
		return
	}

	response, err := services.ModelWrapperService(req)

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"response": response})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
}