package rules

import (
	"errors"
	"receipt-processor-challenge/models"
	"strconv"
)

// Rule which allocates points if total is a round number by implementing the PointsRule interface
type TotalRoundRule struct{}

func (r *TotalRoundRule) GetPoints(receipt *models.Receipt) (int64, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, errors.New("Error parsing Total field")
	}
	floorTotal := float64(int64(total))
	if total-floorTotal == 0.0 {
		return 50, nil
	}
	return 0, nil
}
