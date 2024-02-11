package rules

import (
	"errors"
	"math"
	"receipt-processor-challenge/models"
	"strconv"
	"strings"
)

// Rule which allocates points if item description length is divisible by 3 by implementing the PointsRule interface
type ItemDescriptionLengthRule struct{}

func (r *ItemDescriptionLengthRule) GetPoints(receipt *models.Receipt) (int64, error) {
	points := int64(0)
	for _, item := range receipt.Items {
		trimmedDescLen := len(strings.TrimSpace(item.ShortDescription))
		if trimmedDescLen%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			if err != nil {
				return 0, errors.New("Error parsing Price field")
			}
			points += int64(math.Ceil(price * 0.2))
		}
	}
	return points, nil
}
