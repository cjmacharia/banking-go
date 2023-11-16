package domain

import "time"

type Account struct {
	Id          string    `json:"customer_id"`
	CustomerID  string    `json:"customer_id"`
	AccountType string    `json:"account"`
	Amount      string    `json:"amount"`
	OpeningDate time.Time `json:"opening_date"`
	Status      bool      `json:"status"`
}

type AccountRepository interface {
	CreateAccount(c Account) error
}
