package rules

import (
	"errors"
	"math"
	"receipt-processor-challenge/models"
	"strconv"
)

// Rule which allocates points if total is multiple of 0.25 by implementing the PointsRule interface
type TotalMultipleRule struct{}

func (r *TotalMultipleRule) GetPoints(receipt *models.Receipt) (int64, error) {
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, errors.New("Error parsing Total field")
	}
	if math.Mod(total, 0.25) == 0 {
		return 25, nil
	}
	return 0, nil
}
