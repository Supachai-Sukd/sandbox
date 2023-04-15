//go:build integration

package true_money_wallet

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supachai-sukd/sandboxtesting/config"
)

// Mock ของ
func handlerMock(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"id": 1, "name": "Supachai", "Info": "gopher"}`))
}

func TestMakeHttp(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(handlerMock))
	defer server.Close()

	want := &Response{
		ID:   1,
		Name: "Supachai",
		Info: "gopher",
	}

	t.Run("Happy server response", func(t *testing.T) {
		resp, err := MakeHTTPCall(server.URL)

		// เราสามารถใช้การเทียบ Struct ได้ตามด้านล่าง
		// reflect.DeepEqual คือ ลึกๆแล้ว Struct นั้นหน้าตาเหมือนกันไหม
		if !reflect.DeepEqual(resp, want) {
			t.Errorf("expected (%v), got (%v)", want, resp)
		}

		if !errors.Is(err, nil) {
			t.Errorf("expected (%v), got (%v)", nil, err)
		}
	})

}

func TestCreateWalletHandler(t *testing.T) {
	// Initialize a new gin.Engine instance.
	router := gin.New()

	cfg := config.New().All()
	sql, err := sql.Open("postgres", cfg.DBConnection)
	if err != nil {
		t.Error(err)
	}
	h := New(sql)
	// Use the CreateWallet handler function as the request handler.
	router.POST("/wallets", h.CreateWallet)

	// Create a new http.ResponseWriter to record the response.
	recorder := httptest.NewRecorder()

	// Create a new http.Request with a json payload representing a valid TrueMoneyWallet.
	payload := `{"name": "My Wallet", "category": "General", "currency": "USD", "balance": 1000}`
	request, err := http.NewRequest(http.MethodPost, "/wallets", bytes.NewBufferString(payload))
	if err != nil {
		t.Fatalf("Error creating request: %v", err)
	}

	// Set the content type header to "application/json".
	request.Header.Set("Content-Type", "application/json")

	// Handle the request and record the response.
	router.ServeHTTP(recorder, request)

	// Assert that the response status code is http.StatusCreated.
	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected status code %d, but got %d gggg%s", http.StatusCreated, recorder.Code, recorder.Body.String())
	}

	// Parse the response body as a json object.
	responseBody := make(map[string]interface{})
	err = json.Unmarshal(recorder.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Error parsing response body: %v", err)
	}

	// Assert that the response body matches the expected output format of the TrueMoneyWallet struct.
	expectedOutput := map[string]interface{}{
		"name":     "My Wallet",
		"category": "General",
		"currency": "USD",
		"balance":  1000.0,
	}

	assert.Equal(t, expectedOutput["name"], responseBody["name"])
	assert.Equal(t, expectedOutput["category"], responseBody["category"])
	assert.Equal(t, expectedOutput["currency"], responseBody["currency"])
	assert.Equal(t, expectedOutput["balance"], responseBody["balance"])
}
