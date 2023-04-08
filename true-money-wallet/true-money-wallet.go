package true_money_wallet

import "database/sql"

type TrueMoneyWallet struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Currency string  `json:"currency"`
	Balance  float64 `json:"balance"`
}

type handler struct {
	db *sql.DB
}

func New(db *sql.DB) *handler {
	return &handler{db}
}
