package app

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SaravananPitchaimuthu/pub-api/domain"
	"github.com/SaravananPitchaimuthu/pub-api/logger"
	"github.com/SaravananPitchaimuthu/pub-api/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func Start() {
	router := mux.NewRouter()
	dbClient := getDbClient()
	CustomerRepositoryDb := domain.NewCustomerRepositoryDb(dbClient)
	AccountRepositoryDb := domain.NewAccountRepositoryDb(dbClient)
	ch := CustomerHandler{service.NewCustomerService(CustomerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(AccountRepositoryDb)}

	router.HandleFunc("/customers", ch.getAllCustomer).
		Methods(http.MethodGet).
		Name("GetAllCustomers")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer).
		Methods(http.MethodGet).
		Name("GetCustomer")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).
		Methods(http.MethodPost).
		Name("NewAccount")
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).
		Methods(http.MethodPost).
		Name("NewTransaction")

	address := "127.0.0.1"
	port := "8080"
	logger.Info(fmt.Sprintf("Starting server on %s:%s", address, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {
	dbAddress := "127.0.0.1"
	dbPort := "3306"
	dbUser := "root"
	dbPass := "secret"
	dbName := "Employees"

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbAddress, dbPort, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client

}
