package utils

import (
	"net/http"
	"receipt-processor-challenge/models"

	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, err error) bool {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{Message: err.Error()})
		return true
	}
	return false
}
