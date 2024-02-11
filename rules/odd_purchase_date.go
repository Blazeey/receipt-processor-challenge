package rules

import (
	"errors"
	"receipt-processor-challenge/models"
	"time"
)

// Rule which allocates points if day of the purchase date is odd by implementing the PointsRule interface
type OddPurchaseDateRule struct{}

func (r *OddPurchaseDateRule) GetPoints(receipt *models.Receipt) (int64, error) {
	date, err := time.Parse(time.DateOnly, receipt.PurchaseDate)
	if err != nil {
		return 0, errors.New("Error parsing Purchase Date field")
	}
	return (int64(date.Day()) % 2) * 6, nil
}
