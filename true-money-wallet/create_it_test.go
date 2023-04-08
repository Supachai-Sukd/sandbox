//go:build integration

package true_money_wallet

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateWallet(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define a slice to hold our wallets
	var wallets []TrueMoneyWallet

	// Define a handler function to create a new wallet
	createHandler := func(c *gin.Context) {
		// Parse the request body into a TrueMoneyWallet struct
		var wallet TrueMoneyWallet
		if err := c.BindJSON(&wallet); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Generate a new ID for the wallet
		wallet.ID = len(wallets) + 1

		// Add the new wallet to our slice
		wallets = append(wallets, wallet)

		// Return the created wallet as JSON
		c.JSON(http.StatusCreated, wallet)
	}

	// Mount the createHandler function at the /true-money-wallet endpoint
	router.POST("/true-money-wallet", createHandler)

	// Create a new wallet to send in our request
	newWallet := TrueMoneyWallet{
		Name:     "test1",
		Category: "save",
		Currency: "THB",
		Balance:  99.98,
	}

	// Marshal the new wallet into JSON
	payload, err := json.Marshal(newWallet)
	if err != nil {
		t.Fatalf("Error marshaling payload: %v", err)
	}

	// Create a new request with the JSON payload
	req, err := http.NewRequest("POST", "/true-money-wallet", bytes.NewBuffer(payload))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	// Create a new test HTTP recorder
	recorder := httptest.NewRecorder()

	// Send the request to the server
	router.ServeHTTP(recorder, req)

	// Check that the response code is 201 Created
	assert.Equal(t, http.StatusCreated, recorder.Code)

	// Unmarshal the response body into a TrueMoneyWallet struct
	var createdWallet TrueMoneyWallet
	if err := json.Unmarshal(recorder.Body.Bytes(), &createdWallet); err != nil {
		t.Fatalf("Error unmarshaling response: %v", err)
	}

	// Check that the created wallet has the expected name, category, currency and balance
	assert.Equal(t, newWallet.Name, createdWallet.Name)
	assert.Equal(t, newWallet.Category, createdWallet.Category)
	assert.Equal(t, newWallet.Currency, createdWallet.Currency)
	assert.Equal(t, newWallet.Balance, createdWallet.Balance)

	// Check that the created wallet has a non-zero ID
	assert.NotEqual(t, 0, createdWallet.ID)
}
