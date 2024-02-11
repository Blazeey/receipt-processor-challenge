package handler

import (
	"net/http"
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/storage"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReceiptHandler struct {
	da storage.DataAccess
}

func (h *ReceiptHandler) PostReceiptsProcess(c *gin.Context) {
	var receipt models.Receipt
	if err := c.BindJSON(&receipt); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: err.Error()})
		return
	}
	id := h.da.AddReceipt(&receipt)
	c.JSON(http.StatusOK, models.ReceiptCreationResponse{Id: id.String()})
}

func (h *ReceiptHandler) GetReceiptsIdPoints(c *gin.Context) {
	id := c.Param("id")
	receiptId, err := uuid.Parse(id)
	if Error(c, err) {
		return
	}
	points, err := h.da.GetPoints(receiptId)
	if Error(c, err) {
		return
	}
	c.JSON(http.StatusOK, models.GetPointsResponse{Points: points})
}

func GetCachedReceiptHandler() ReceiptHandler {
	return ReceiptHandler{
		da: storage.GetCachedAccess(&storage.InMemoryStore{}),
	}
}

func Error(c *gin.Context, err error) bool {
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(http.StatusNotFound, models.ErrorResponse{Message: err.Error()})
		return true
	}
	return false
}
