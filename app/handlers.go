package app

import (
	"encoding/json"
	"fmt"
	"github.com/cjmash/banking/domain"
	"github.com/cjmash/banking/domain/service"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}
type CustomerHandler struct {
	service service.CustomerService
}

type AccountHandler struct {
	service service.AccountService
}

func (ch *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandler) PostCustomer(w http.ResponseWriter, r *http.Request) {
	c := &domain.Customer{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = ch.service.PostCustomer(c)
	fmt.Println(err)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(c)
}

func (ah *AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	c := &domain.Account{}
	err := json.NewDecoder(r.Body).Decode(c)
	if err != nil {
		fmt.Println(err, "hehrh")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	err = ah.service.CreateAccount(*c)
	fmt.Println(err)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(c)
}
