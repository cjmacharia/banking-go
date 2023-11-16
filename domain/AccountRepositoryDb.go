package domain

import (
	"context"
	"fmt"
	"time"
)

func (db RepositoryDB) CreateAccount(a Account) error {
	sql := "INSERT into accounts (customer_id,account_type, amount, opening_date, status)VALUES ($1, $2, $3, $4, $5) "
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := db.Client.ExecContext(ctx, sql, a.CustomerID, a.AccountType, a.Amount, a.OpeningDate, a.Status)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
