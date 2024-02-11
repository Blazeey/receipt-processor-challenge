package rules

import (
	"receipt-processor-challenge/models"
	"unicode"
)

// Rule which allocates points based on number of word and digit characters by implementing the PointsRule interface
type AlphaNumRule struct{}

func (r *AlphaNumRule) GetPoints(receipt *models.Receipt) (int64, error) {
	count := 0
	for _, c := range receipt.Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			count++
		}
	}
	return int64(count), nil
}
