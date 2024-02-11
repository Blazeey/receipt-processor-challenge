package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"receipt-processor-challenge/models"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPostReceiptsProcess(t *testing.T) {
	s := NewServer("localhost", 9090)
	requestBody, _ := json.Marshal(getTargetReceipt())
	request, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer(requestBody))
	response := httptest.NewRecorder()
	s.router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "Expected HTTP 200 OK, got: %v", response.Code)
}

func TestPostReceiptsProcessInvalid(t *testing.T) {

	invalidRetailerReceipt := getTargetReceipt()
	invalidRetailerReceipt.Retailer = ""
	testInvalidReceipt(t, invalidRetailerReceipt)

	invalidPurchaseDateReceipt := getTargetReceipt()
	invalidPurchaseDateReceipt.PurchaseDate = "2022-1-01"
	testInvalidReceipt(t, invalidPurchaseDateReceipt)

	invalidPurchaseTimeReceipt := getTargetReceipt()
	invalidPurchaseTimeReceipt.PurchaseTime = "23:1"
	testInvalidReceipt(t, invalidPurchaseTimeReceipt)

	invalidTotalReceipt := getTargetReceipt()
	invalidTotalReceipt.Total = "23.1"
	testInvalidReceipt(t, invalidTotalReceipt)

	invalidItemPriceReceipt := getTargetReceipt()
	invalidItemPriceReceipt.Items[0].Price = "23.1"
	testInvalidReceipt(t, invalidItemPriceReceipt)

	invalidShortDescReceipt := getTargetReceipt()
	invalidShortDescReceipt.Items[0].ShortDescription = "asdasdasd&"
	testInvalidReceipt(t, invalidShortDescReceipt)
}

func TestGetReceiptsIdPoints(t *testing.T) {
	s := NewServer("localhost", 9090)
	requestBody, _ := json.Marshal(getTargetReceipt())
	request, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer(requestBody))
	response := httptest.NewRecorder()
	s.router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "Expected HTTP 200 OK, got: %v", response.Code)
	var responseBody models.ReceiptCreationResponse
	defer response.Result().Body.Close()
	decorder := json.NewDecoder(response.Body)
	err := decorder.Decode(&responseBody)
	assert.Nil(t, err)

	pointsRequest, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/receipts/%s/points", responseBody.Id), &bytes.Buffer{})
	pointsResponse := httptest.NewRecorder()
	s.router.ServeHTTP(pointsResponse, pointsRequest)
	assert.Equal(t, http.StatusOK, response.Code)
	var pointsResponseBody models.GetPointsResponse
	defer pointsResponse.Result().Body.Close()
	pointsDecoder := json.NewDecoder(pointsResponse.Body)
	err = pointsDecoder.Decode(&pointsResponseBody)
	assert.Nil(t, err)
	assert.Equal(t, int64(2), pointsResponseBody.Points)
}

func TestReceiptNotFoundForPoints(t *testing.T) {
	s := NewServer("localhost", 9090)
	requestBody, _ := json.Marshal(getTargetReceipt())
	request, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer(requestBody))
	response := httptest.NewRecorder()
	s.router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code, "Expected HTTP 200 OK, got: %v", response.Code)
	var responseBody models.ReceiptCreationResponse
	defer response.Result().Body.Close()
	decorder := json.NewDecoder(response.Body)
	err := decorder.Decode(&responseBody)
	assert.Nil(t, err)

	pointsRequest, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/receipts/%s/points", uuid.New()), &bytes.Buffer{})
	pointsResponse := httptest.NewRecorder()
	s.router.ServeHTTP(pointsResponse, pointsRequest)
	assert.Equal(t, http.StatusNotFound, pointsResponse.Code)
	var pointsResponseBody models.GetPointsResponse
	defer pointsResponse.Result().Body.Close()
	pointsDecoder := json.NewDecoder(pointsResponse.Body)
	err = pointsDecoder.Decode(&pointsResponseBody)
	assert.Nil(t, err)
}

func testInvalidReceipt(t *testing.T, receipt models.Receipt) {
	s := NewServer("localhost", 9090)
	requestBody, _ := json.Marshal(receipt)
	request, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer(requestBody))
	response := httptest.NewRecorder()
	s.router.ServeHTTP(response, request)
	assert.Equal(t, http.StatusBadRequest, response.Code, "Expected HTTP 400 Bad Request, got: %v", response.Code)
}

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
