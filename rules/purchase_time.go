package rules

import (
	"errors"
	"receipt-processor-challenge/models"
	"time"
)

// Rule which allocates points if purchase time is after 2pm and before 4pm by implementing the PointsRule interface
type PurchaseTimeRule struct{}

func (r *PurchaseTimeRule) GetPoints(receipt *models.Receipt) (int64, error) {
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, errors.New("Error parsing Purchase Time field")
	}

	minimumTime, _ := time.Parse("15:04", "14:00")
	maximumTime, _ := time.Parse("15:04", "16:00")

	if purchaseTime.After(minimumTime) && purchaseTime.Before(maximumTime) {
		return 10, nil
	}
	return 0, nil
}
