package storage

import (
	"receipt-processor-challenge/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMemoryAccess(t *testing.T) {
	store := &InMemoryStore{}
	receipt := getTargetReceipt()
	id := store.AddReceipt(&receipt)
	val, ok := store.data.Load(id)
	assert.True(t, ok)
	assert.Equal(t, receipt, val.(models.Receipt))

	points, err := store.GetPoints(id)
	assert.Nil(t, err)
	assert.Equal(t, int64(28), points)
}

func TestMemoryAccessReceiptNotFound(t *testing.T) {
	store := &InMemoryStore{}
	receipt := getTargetReceipt()
	_ = store.AddReceipt(&receipt)
	_, ok := store.data.Load(uuid.New())
	assert.False(t, ok)

	_, err := store.GetPoints(uuid.New())
	assert.NotNil(t, err)
}

func getTargetReceipt() models.Receipt {
	return models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []*models.Item{
			&models.Item{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			},
			&models.Item{
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			},
			&models.Item{
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			},
			&models.Item{
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			},
			&models.Item{
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            "12.00",
			},
		},
	}
}
