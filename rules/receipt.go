package rules

import "receipt-processor-challenge/models"

// Rule which allocates points by containing multiple PointsRule
type ReceiptPointsRule struct {
	rules []PointsRule
}

func (r *ReceiptPointsRule) GetPoints(receipt *models.Receipt) (int64, error) {
	points := int64(0)
	for _, rule := range r.rules {
		point, err := rule.GetPoints(receipt)
		if err != nil {
			return 0, err
		}
		points += point
	}
	return points, nil
}
