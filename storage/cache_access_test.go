package storage

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCacheAccess(t *testing.T) {
	store := &CachedAccess{
		db: &InMemoryStore{},
	}
	receipt := getTargetReceipt()
	id := store.AddReceipt(&receipt)
	_, ok := store.data.Load(id)
	assert.False(t, ok)

	points, err := store.GetPoints(id)
	assert.Nil(t, err)
	assert.Equal(t, int64(28), points)
	cachedPoints, ok := store.data.Load(id)
	assert.True(t, ok)
	assert.Equal(t, int64(28), cachedPoints)
}

func TestCacheAccessReceiptNotFound(t *testing.T) {
	store := &CachedAccess{
		db: &InMemoryStore{},
	}
	receipt := getTargetReceipt()
	_ = store.AddReceipt(&receipt)
	_, ok := store.data.Load(uuid.New())
	assert.False(t, ok)

	_, err := store.GetPoints(uuid.New())
	assert.NotNil(t, err)
}
