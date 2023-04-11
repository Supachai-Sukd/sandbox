//go:build unit

package true_money_wallet




import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"bytes"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateWallet(t *testing.T) {
	// Set up Gin
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	// Set up mock DB
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error initializing mock database: %s", err)
	}
	defer db.Close()

	h := handler{db: db}
	wallet := &TrueMoneyWallet{
		Name:     "Test Wallet",
		Category: "Test Category",
		Currency: "USD",
		Balance:  100,
	}

	// Set up mock expectations
	mock.ExpectPrepare("INSERT INTO true_money_wallet").ExpectQuery().
		WithArgs(wallet.Name, wallet.Category, wallet.Currency, wallet.Balance).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Create request
	walletJSON, _ := json.Marshal(wallet)
	req, _ := http.NewRequest("POST", "/wallet", bytes.NewBuffer(walletJSON))
	req.Header.Set("Content-Type", "application/json")

	// Execute request
	resp := httptest.NewRecorder()
	router.POST("/wallet", h.CreateWallet)
	router.ServeHTTP(resp, req)

	// Check response
	assert.Equal(t, http.StatusCreated, resp.Code)

	// Check response body
	var respWallet TrueMoneyWallet
	err = json.Unmarshal(resp.Body.Bytes(), &respWallet)
	assert.Nil(t, err)
	assert.Equal(t, "Test Wallet", respWallet.Name)
	assert.Equal(t, "Test Category", respWallet.Category)
	assert.Equal(t, "USD", respWallet.Currency)
	assert.Equal(t, 100.00, respWallet.Balance)

	// Check mock expectations were met
	err = mock.ExpectationsWereMet()
	assert.Nil(t, err)
}
