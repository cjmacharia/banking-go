package domain

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

// FindAll adaptor for the secondary port
func (db RepositoryDB) FindAll() ([]Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	customerSQL := "select customer_id, name, city, zipcode, date_of_birth, status, created_at from customers"
	rows, err := db.Client.QueryContext(ctx, customerSQL)

	if err != nil {
		fmt.Println("f:", err)
		return nil, err
	}
	var customers []Customer
	for rows.Next() {
		var customer Customer
		err := rows.Scan(
			&customer.Id,
			&customer.Name,
			&customer.City,
			&customer.Zipcode,
			&customer.DateOfBirth,
			&customer.Status,
			&customer.CreatedAt,
		)
		if err != nil {
			fmt.Println("Error scanning row:", err)
			return customers, err
		}
		customers = append(customers, customer)
	}
	return customers, nil

}

func (db RepositoryDB) CreateCustomer(c *Customer) error {
	stmt := "insert into customers (name, city, zipcode, date_of_birth, status, created_at) VALUES ($1, $2, $3, $4, $5, $6) "
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := db.Client.ExecContext(ctx, stmt, c.Name, c.City, c.Zipcode, c.DateOfBirth, c.Status, c.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func NewCustomerRepositoryDB(dbClient *sql.DB) RepositoryDB {
	return RepositoryDB{Client: dbClient}
}
