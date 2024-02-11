package server

import (
	"receipt-processor-challenge/models"

	"github.com/gin-gonic/gin"
)

// Initializes heartbeat route
func InitializeHeartbeatRoute(router *gin.RouterGroup) {
	router.GET("/heartbeat", heartbeat)
}

func heartbeat(c *gin.Context) {
	c.JSON(200, models.DefaultSuccessResponse{
		Message: "OK",
	})
}
