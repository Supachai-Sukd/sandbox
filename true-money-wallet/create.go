package true_money_wallet

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h handler) CreateWallet(c *gin.Context) {
	wallet := new(TrueMoneyWallet)
	if err := c.Bind(wallet); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	stmt, err := h.db.Prepare("INSERT INTO true_money_wallet (name, category, currency, balance) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(wallet.Name, wallet.Category, wallet.Currency, wallet.Balance)
	err = row.Scan(&wallet.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusCreated, TrueMoneyWallet{
		ID:       wallet.ID,
		Name:     wallet.Name,
		Category: wallet.Category,
		Currency: wallet.Currency,
		Balance:  wallet.Balance,
	})
}



type Response struct { 
	ID          int    `json:"id"` 
	Name        string `json:"name"` 
	Info 		string `json:"info"`
}


var get = http.Get

func MakeHTTPCall(url string) (*Response, error) { 
	resp, err := http.Get(url) 
	if err != nil { 
		return nil, err 
	} 

		body, err := ioutil.ReadAll(resp.Body) 

		if err != nil {
			 return nil, err 
		} 

	r := &Response{} 

	if err := json.Unmarshal(body, r); err !=  nil {
		return nil, err
	}
		return r, nil
}	