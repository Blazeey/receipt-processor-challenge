package rules

import "receipt-processor-challenge/models"

// Rule which allocates points based on number of pairs of items found by implementing the PointsRule interface
type PairItemsRule struct{}

func (r *PairItemsRule) GetPoints(receipt *models.Receipt) (int64, error) {
	return int64((len(receipt.Items) / 2) * 5), nil
}
