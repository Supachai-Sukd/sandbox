package true_money_wallet

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) WalletList(c *gin.Context) {
	// Query the database for all wallets
	rows, err := h.db.Query("SELECT * FROM true_money_wallet")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()

	// Parse the rows into wallet structs
	var wallets []TrueMoneyWallet
	for rows.Next() {
		var wallet TrueMoneyWallet
		err := rows.Scan(&wallet.ID, &wallet.Name, &wallet.Category, &wallet.Currency, &wallet.Balance)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		wallets = append(wallets, wallet)
	}

	// Return the wallets as a JSON response
	c.JSON(http.StatusOK, wallets)
}

