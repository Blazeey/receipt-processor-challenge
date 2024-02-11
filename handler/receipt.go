package handler

import (
	"errors"
	"fmt"
	"net/http"
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/storage"
	"receipt-processor-challenge/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ReceiptHandler struct {
	da storage.DataAccess
}

// Handler for POST /receipts/process
func (h *ReceiptHandler) PostReceiptsProcess(c *gin.Context) {
	var receipt models.Receipt
	if err := c.BindJSON(&receipt); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: fmt.Sprintf("Field %s is either not present or invalid", ve[0].Field())})
		} else {
			c.JSON(http.StatusBadRequest, models.ErrorResponse{Message: "The receipt is invalid"})
		}
		return
	}
	id := h.da.AddReceipt(&receipt)
	c.JSON(http.StatusOK, models.ReceiptCreationResponse{Id: id.String()})
}

// Handler for GET /receipts/:id/points
func (h *ReceiptHandler) GetReceiptsIdPoints(c *gin.Context) {
	id := c.Param("id")
	receiptId, err := uuid.Parse(id)
	if utils.Error(c, err) {
		return
	}
	points, err := h.da.GetPoints(receiptId)
	if utils.Error(c, err) {
		return
	}
	c.JSON(http.StatusOK, models.GetPointsResponse{Points: points})
}

// Handler to be registered in routes
func GetCachedReceiptHandler() ReceiptHandler {
	return ReceiptHandler{
		da: storage.GetCachedAccess(new(storage.InMemoryStore)),
	}
}
