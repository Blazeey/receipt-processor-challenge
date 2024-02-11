package storage

import (
	"errors"
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/rules"
	"sync"

	"github.com/google/uuid"
)

var defaultPointsRule = rules.GetDefaultReceiptPointsRule()

// This stores the actual data in memory. Its easier to swap the underlying storage later.
// If we are connecting to a database in the future, we can swap this out with a new struct implementing
// Access interface
type InMemoryStore struct {
	data sync.Map
}

func (s *InMemoryStore) AddReceipt(receipt *models.Receipt) uuid.UUID {
	id := uuid.New()
	s.data.Store(id, *receipt)
	return id
}

func (s *InMemoryStore) GetPoints(receiptId uuid.UUID) (int64, error) {
	value, found := s.data.Load(receiptId)
	if !found {
		return 0, errors.New("Receipt not found")
	}
	receipt := value.(models.Receipt)
	return defaultPointsRule.GetPoints(&receipt)
}
