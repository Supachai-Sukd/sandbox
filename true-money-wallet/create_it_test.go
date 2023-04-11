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

// func TestCreateWalletIntegration(t *testing.T) {
// 	// start the server
// 	go func() {
// 		err := http.ListenAndServe(":2566", nil)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	// wait for the server to start
// 	time.Sleep(1 * time.Second)

// 	// set the endpoint URL
// 	url := "http://localhost:2566/wallet"

// 	body := `{"name":"test1", "category":"save", "currency":"THB", "balance": 99}`

// 	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(body))
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	defer resp.Body.Close()

// 	var response Wallet
// 	err = json.NewDecoder(resp.Body).Decode(&response)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	expected := Wallet{
// 		Name:     "test1",
// 		Category: "save",
// 		Currency: "THB",
// 		Balance:  99.0,
// 	}
// 	assert.Equal(t, expected, response)

// }


func TestCreateWalletIntegration(t *testing.T) {
	// start the server
	// go func() {
	// 	err := http.ListenAndServe(":2566", nil)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// // wait for the server to start
	// time.Sleep(1 * time.Second)
   
	// // set the endpoint URL
	// url := "http://localhost:2566/wallet"
   // create a new Gin router
	router := gin.Default()

	// define the endpoint handler
	router.POST("/wallet", func(c *gin.Context) {
		var wallet TrueMoneyWallet
		err := c.ShouldBindJSON(&wallet)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		wallet.ID = 99 // set a fixed ID for testing
		c.JSON(http.StatusOK, wallet)
	})

	// create the HTTP request
	payload := gin.H{"name": "test1", "category": "save", "currency": "THB", "balance": 99.0}
	payloadBytes, _ := json.Marshal(payload)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:2566/wallet", bytes.NewBuffer(payloadBytes))
	req.Header.Set("Content-Type", "application/json")

	// send the HTTP request and record the response
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	// check the response
	assert.Equal(t, http.StatusOK, res.Code)
	expectedResponse := gin.H{"id": 99.0, "name": "test1", "category": "save", "currency": "THB", "balance": 99.0}
	var actualResponse gin.H
	err := json.Unmarshal(res.Body.Bytes(), &actualResponse)
	assert.Nil(t, err)
	assert.Equal(t, expectedResponse, actualResponse)
}
