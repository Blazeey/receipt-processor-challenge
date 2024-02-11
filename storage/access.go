package storage

import (
	"receipt-processor-challenge/models"

	"github.com/google/uuid"
)

type DataAccess interface {
	AddReceipt(receipt *models.Receipt) uuid.UUID
	GetPoints(receiptId uuid.UUID) (int64, error)
}
