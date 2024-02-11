package storage

import (
	"receipt-processor-challenge/models"
	"sync"

	"github.com/google/uuid"
)

type CachedAccess struct {
	db   DataAccess
	data sync.Map
}

func (s *CachedAccess) AddReceipt(receipt *models.Receipt) uuid.UUID {
	return s.db.AddReceipt(receipt)
}

func (s *CachedAccess) GetPoints(receiptId uuid.UUID) (int64, error) {
	value, found := s.data.Load(receiptId)
	if found {
		return value.(int64), nil
	}
	points, err := s.db.GetPoints(receiptId)
	if err != nil {
		return 0, err
	}
	s.data.Store(receiptId, points)
	return points, nil
}

func GetCachedAccess(da DataAccess) DataAccess {
	return &CachedAccess{
		db: da,
	}
}
