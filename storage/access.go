package storage

import (
	"receipt-processor-challenge/models"

	"github.com/google/uuid"
)

// This interface is used to fetch the data form in memory DB or cache.
type DataAccess interface {
	AddReceipt(receipt *models.Receipt) uuid.UUID
	GetPoints(receiptId uuid.UUID) (int64, error)
}
