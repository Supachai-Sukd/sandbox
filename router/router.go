package router

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"

	true_money_wallet "github.com/supachai-sukd/sandboxtesting/true-money-wallet"
)

func RegRoute(db *sql.DB) *gin.Engine {
	r := gin.New()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	hWallet := true_money_wallet.New(db)

	r.GET("/wallets", hWallet.WalletList)
	r.POST("/wallet", hWallet.CreateWallet)

	return r
}
