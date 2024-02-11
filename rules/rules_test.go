package rules

import (
	"receipt-processor-challenge/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getTargetReceipt() models.Receipt {
	return models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Total:        "35.35",
		Items: []*models.Item{
			&models.Item{
				ShortDescription: "Mountain Dew 12PK",
				Price:            "6.49",
			},
			&models.Item{
				ShortDescription: "Emils Cheese Pizza",
				Price:            "12.25",
			},
			&models.Item{
				ShortDescription: "Knorr Creamy Chicken",
				Price:            "1.26",
			},
			&models.Item{
				ShortDescription: "Doritos Nacho Cheese",
				Price:            "3.35",
			},
			&models.Item{
				ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
				Price:            "12.00",
			},
		},
	}
}

func getMnMReceipt() models.Receipt {
	return models.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Total:        "9.00",
		Items: []*models.Item{
			&models.Item{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			&models.Item{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			&models.Item{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
			&models.Item{
				ShortDescription: "Gatorade",
				Price:            "2.25",
			},
		},
	}
}

func TestAlphaNumRule(t *testing.T) {
	rule := AlphaNumRule{}

	targetReceipt := getTargetReceipt()
	targetPoints, err := rule.GetPoints(&targetReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(6), targetPoints, "Expected 6, got %d", targetPoints)

	mnmReceipt := getMnMReceipt()
	mnmPoints, err := rule.GetPoints(&mnmReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(14), mnmPoints, "Expected 14, got %d", targetPoints)
}

func TestTotalRoundRule(t *testing.T) {
	rule := TotalRoundRule{}

	targetReceipt := getTargetReceipt()
	targetPoints, err := rule.GetPoints(&targetReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), targetPoints, "Expected 0, got %d", targetPoints)

	mnmReceipt := getMnMReceipt()
	mnmPoints, err := rule.GetPoints(&mnmReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(50), mnmPoints, "Expected 50, got %d", targetPoints)
}

func TestTotalMultipleRule(t *testing.T) {
	rule := TotalMultipleRule{}

	targetReceipt := getTargetReceipt()
	targetPoints, err := rule.GetPoints(&targetReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), targetPoints, "Expected 0, got %d", targetPoints)

	mnmReceipt := getMnMReceipt()
	mnmPoints, err := rule.GetPoints(&mnmReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(25), mnmPoints, "Expected 25, got %d", targetPoints)
}

func TestPairItemsRule(t *testing.T) {
	rule := PairItemsRule{}

	targetReceipt := getTargetReceipt()
	targetPoints, err := rule.GetPoints(&targetReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), targetPoints, "Expected 10, got %d", targetPoints)

	mnmReceipt := getMnMReceipt()
	mnmPoints, err := rule.GetPoints(&mnmReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), mnmPoints, "Expected 10, got %d", targetPoints)
}

func TestItemDescriptionLengthRule(t *testing.T) {
	rule := ItemDescriptionLengthRule{}

	targetReceipt := getTargetReceipt()
	targetPoints, err := rule.GetPoints(&targetReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(6), targetPoints, "Expected 6, got %d", targetPoints)

	mnmReceipt := getMnMReceipt()
	mnmPoints, err := rule.GetPoints(&mnmReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), mnmPoints, "Expected 0, got %d", targetPoints)
}

func TestOddPurchaseDateRule(t *testing.T) {
	rule := OddPurchaseDateRule{}

	targetReceipt := getTargetReceipt()
	targetPoints, err := rule.GetPoints(&targetReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(6), targetPoints, "Expected 6, got %d", targetPoints)

	mnmReceipt := getMnMReceipt()
	mnmPoints, err := rule.GetPoints(&mnmReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), mnmPoints, "Expected 0, got %d", targetPoints)
}

func TestPurchaseTimeRule(t *testing.T) {
	rule := PurchaseTimeRule{}

	targetReceipt := getTargetReceipt()
	targetPoints, err := rule.GetPoints(&targetReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), targetPoints, "Expected 0, got %d", targetPoints)

	mnmReceipt := getMnMReceipt()
	mnmPoints, err := rule.GetPoints(&mnmReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), mnmPoints, "Expected 10, got %d", targetPoints)
}

func TestReceiptPointsRule(t *testing.T) {
	rule := GetDefaultReceiptPointsRule()

	targetReceipt := getTargetReceipt()
	targetPoints, err := rule.GetPoints(&targetReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(28), targetPoints, "Expected 28, got %d", targetPoints)

	mnmReceipt := getMnMReceipt()
	mnmPoints, err := rule.GetPoints(&mnmReceipt)
	assert.NoError(t, err)
	assert.Equal(t, int64(109), mnmPoints, "Expected 109, got %d", targetPoints)
}
