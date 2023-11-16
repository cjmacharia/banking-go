package app

import (
	"fmt"
	"github.com/cjmash/banking/domain"
	"github.com/cjmash/banking/domain/service"
	mux "github.com/gorilla/mux"
	"net/http"
)

func Start() {
	dbClient := domain.GetDbClient()
	ch := CustomerHandler{service.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient))}
	ah := AccountHandler{service.NewAccountService(domain.NewCustomerRepositoryDB(dbClient))}

	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.PostCustomer).Methods(http.MethodPost)
	router.HandleFunc("/account", ah.CreateAccount).Methods(http.MethodPost)

	err := http.ListenAndServe("localhost:8080", router)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}
}
