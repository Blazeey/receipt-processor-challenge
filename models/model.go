package models

// Item defines model for Item.
type Item struct {
	// Price The total price payed for this item.
	Price string `json:"price" binding:"required,regex=^\\d+\\.\\d{2}$"`

	// ShortDescription The Short Product Description for the item.
	ShortDescription string `json:"shortDescription" binding:"required,regex=^[\\w\\s\\-]+$"`
}

// Receipt defines model for Receipt.
type Receipt struct {
	Items []*Item `json:"items" binding:"required,dive"`

	// PurchaseDate The date of the purchase printed on the receipt.
	PurchaseDate string `json:"purchaseDate" binding:"required,time=2006-01-02"`

	// PurchaseTime The time of the purchase printed on the receipt. 24-hour time expected.
	PurchaseTime string `json:"purchaseTime" binding:"required,time=15:04"`

	// Retailer The name of the retailer or store the receipt is from.
	Retailer string `json:"retailer" binding:"required,regex=^[\\w\\s\\-&]+$"`

	// Total The total amount paid on the receipt.
	Total string `json:"total" binding:"required,regex=^\\d+\\.\\d{2}$"`
}

type ReceiptCreationResponse struct {
	Id string `json:"id"`
}

type GetPointsResponse struct {
	Points int64 `json:"points"`
}

type DefaultSuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
	//TODO: Add error code?
}
