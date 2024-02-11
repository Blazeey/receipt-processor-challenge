package server

import (
	"receipt-processor-challenge/handler"

	"github.com/gin-gonic/gin"
)

func InitializeReceiptRoutes(router *gin.RouterGroup) {
	receiptsHandler := handler.GetCachedReceiptHandler()
	receiptsRouter := router.Group("/receipts")
	receiptsRouter.POST("/process", receiptsHandler.PostReceiptsProcess)
	receiptsRouter.GET("/:id/points", receiptsHandler.GetReceiptsIdPoints)
}
