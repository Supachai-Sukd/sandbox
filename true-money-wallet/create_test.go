//go:build unit

package true_money_wallet

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateWallet(t *testing.T) {
	// Create a new Gin router
	router := gin.Default()

	// Define a handler function to create a new wallet
	createHandler := func(c *gin.Context) {
		// Parse the request body into a TrueMoneyWallet struct
		var wallet TrueMoneyWallet
		if err := c.BindJSON(&wallet); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Open a new database connection
		db, mock, err := sqlmock.New()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer db.Close()

		// Expect a query to insert the new wallet into the database
		mock.ExpectExec("INSERT INTO true_money_wallet").WithArgs(wallet.Name, wallet.Category, wallet.Currency, wallet.Balance).WillReturnResult(sqlmock.NewResult(1, 1))

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
}
