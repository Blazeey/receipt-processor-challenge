package rules

import (
	"errors"
	"math"
	"receipt-processor-challenge/models"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type PointsRule interface {
	GetPoints(receipt *models.Receipt) (int64, error)
}

type AlphaNumRule struct{}
type TotalRoundRule struct{}
type TotalMultipleRule struct{}
type PairItemsRule struct{}
type ItemDescriptionLengthRule struct{}
type OddPurchaseDateRule struct{}
type PurchaseTimeRule struct{}
type ReceiptPointsRule struct {
	rules []PointsRule
}

func (r *AlphaNumRule) GetPoints(receipt *models.Receipt) (int64, error) {
	count := 0
	for _, c := range receipt.Retailer {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			count++
		}
	}
	return int64(count), nil
}

func (r *TotalRoundRule) GetPoints(receipt *models.Receipt) (int64, error) {
	//TODO: Avoid multiple float parsing?
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

func (r *TotalMultipleRule) GetPoints(receipt *models.Receipt) (int64, error) {
	//TODO: Avoid multiple float parsing?
	total, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0, errors.New("Error parsing Total field")
	}
	if math.Mod(total, 0.25) == 0 {
		return 25, nil
	}
	return 0, nil
}

func (r *PairItemsRule) GetPoints(receipt *models.Receipt) (int64, error) {
	return int64((len(receipt.Items) / 2) * 5), nil
}

func (r *ItemDescriptionLengthRule) GetPoints(receipt *models.Receipt) (int64, error) {
	points := int64(0)
	for _, item := range receipt.Items {
		trimmedDescLen := len(strings.TrimSpace(item.ShortDescription))
		if trimmedDescLen%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
			//TODO: Avoid multiple float parsing?
			if err != nil {
				return 0, errors.New("Error parsing Price field")
			}
			points += int64(math.Ceil(price * 0.2))
		}
	}
	return points, nil
}

func (r *OddPurchaseDateRule) GetPoints(receipt *models.Receipt) (int64, error) {
	date, err := time.Parse(time.DateOnly, receipt.PurchaseDate)
	//TODO: handle error
	if err != nil {
		return 0, errors.New("Error parsing Purchase Date field")
	}
	return (int64(date.Day()) % 2) * 6, nil
}

func (r *PurchaseTimeRule) GetPoints(receipt *models.Receipt) (int64, error) {
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	//TODO: handle error
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
