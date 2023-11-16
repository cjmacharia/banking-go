package domain

import "time"

type Customer struct {
	Id          string    `json:"customer_id"`
	Name        string    `json:"name"`
	City        string    `json:"city"`
	Zipcode     string    `json:"zipcode"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
}

// CustomerRepository secondary port
type CustomerRepository interface {
	FindAll() ([]Customer, error)
	CreateCustomer(c *Customer) error
}
