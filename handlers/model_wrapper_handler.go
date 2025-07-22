package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"pokerdegen/services"
	"pokerdegen/models"
)

func ModelWrapperHandler(c *gin.Context) {
	var req models.ModelRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response": "invalid input to model wrapper service",
			"error": err.Error(),
		})
		return
	}

	response, err := services.ModelWrapperService(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"response": "failed to call model wrapper",
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"response": response,
		})
	}
}