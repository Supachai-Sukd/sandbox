//go:build unit

package true_money_wallet


import (
	"net/http/httptest"
	"testing"
	"encoding/json"
	"net/http"
	"reflect"
	

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

func TestWalletList(t *testing.T) {
	// Set up a mock database
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("error setting up mock database: %s", err)
	}
	defer mockDB.Close()

	// Define the expected results
	expected := []TrueMoneyWallet{
		{
			ID:       1,
			Name:     "test-name",
			Category: "test-category",
			Currency: "test-currency",
			Balance:  9.1,
		},
		{
			ID:       2,
			Name:     "test1",
			Category: "save",
			Currency: "THB",
			Balance:  99.98,
		},
		// Add additional expected results as needed
	}

	// Set up the mock query result
	columns := []string{"id", "name", "category", "currency", "balance"}
	rows := sqlmock.NewRows(columns)
	for _, wallet := range expected {
		rows = rows.AddRow(wallet.ID, wallet.Name, wallet.Category, wallet.Currency, wallet.Balance)
	}
	mock.ExpectQuery("SELECT \\* FROM true_money_wallet").WillReturnRows(rows)

	// Set up the request context
	req, err := http.NewRequest(http.MethodGet, "/wallets", nil)
	if err != nil {
		t.Fatalf("error setting up request: %s", err)
	}
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Set up the handler and call the function
	h := handler{db: mockDB}
	h.WalletList(c)

	// Check the response code
	if w.Code != http.StatusOK {
		t.Errorf("expected response code %d, got %d", http.StatusOK, w.Code)
	}

	// Check the response body
	var actual []TrueMoneyWallet
	if err := json.Unmarshal(w.Body.Bytes(), &actual); err != nil {
		t.Fatalf("error unmarshalling response body: %s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected response body %+v, got %+v", expected, actual)
	}
}

