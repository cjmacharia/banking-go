package domain

import "time"

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "OO1", Name: "ASHISH", City: "Nairobi", Zipcode: "1001", DateOfBirth: time.Date(2001, time.September, 9, 0, 0, 0, 0, time.UTC), Status: true},
	}
	return CustomerRepositoryStub{customers}
}
