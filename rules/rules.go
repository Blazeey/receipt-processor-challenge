package rules

import (
	"receipt-processor-challenge/models"
)

// Points are allocated based on each rule. This PointsRule interface defines the points allocated for a particular rule
type PointsRule interface {
	GetPoints(receipt *models.Receipt) (int64, error)
}

// Default rule contains all the 7 different rules
func GetDefaultReceiptPointsRule() PointsRule {
	defaultRule := ReceiptPointsRule{
		rules: []PointsRule{
			&AlphaNumRule{},
			&TotalRoundRule{},
			&TotalMultipleRule{},
			&PairItemsRule{},
			&ItemDescriptionLengthRule{},
			&OddPurchaseDateRule{},
			&PurchaseTimeRule{}},
	}
	return &defaultRule
}
