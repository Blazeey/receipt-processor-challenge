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
	Items []*Item `json:"items" binding:"required,gt=0,dive"`

	// PurchaseDate The date of the purchase printed on the receipt.
	PurchaseDate string `json:"purchaseDate" binding:"required,time=2006-01-02"`

	// PurchaseTime The time of the purchase printed on the receipt. 24-hour time expected.
	PurchaseTime string `json:"purchaseTime" binding:"required,time=15:04"`

	// Retailer The name of the retailer or store the receipt is from.
	Retailer string `json:"retailer" binding:"required,gt=0"`

	// Total The total amount paid on the receipt.
	Total string `json:"total" binding:"required,regex=^\\d+\\.\\d{2}$"`
}

// Receipt process API response
type ReceiptCreationResponse struct {
	Id string `json:"id"`
}

// Calcualte points API response
type GetPointsResponse struct {
	Points int64 `json:"points"`
}

// Default API response
type DefaultSuccessResponse struct {
	Message string `json:"message"`
}

// Error API response
type ErrorResponse struct {
	Message string `json:"message"`
	//TODO: Add error code?
}
