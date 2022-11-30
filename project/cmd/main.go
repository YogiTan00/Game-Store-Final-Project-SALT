package main

import (
	"context"
	"fmt"
	"game-store-final-project/project/internal/delivery/http/customer_handler"
	"game-store-final-project/project/internal/delivery/http/transaction_handler"
	repo "game-store-final-project/project/internal/repository/mysql/customer"
	"game-store-final-project/project/internal/repository/mysql/transaction"
	"game-store-final-project/project/internal/usecase/customer"
	"game-store-final-project/project/pkg/mysql_connection"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ctx                   = context.Background()
	mysqlConn             = mysql_connection.InitMysqlDB()
	repoCustomer          = repo.NewCustomerRepositoryMysqlInteractor(mysqlConn)
	useCaseCustomer       = customer.NewCustomerUseCaseInteractor(ctx, repoCustomer)
	repoTransaction       = transaction.NewTransactionMysqlInteractor(mysqlConn)
	repoTransactionDetail = transaction.NewTransactionDetailMysqlInteractor(mysqlConn)
)

func main() {
	r := mux.NewRouter()

	// routes customer
	handlerCustomer := customer_handler.NewCustomerHandler(useCaseCustomer)
	handlerTransaction := transaction_handler.NewTransactionHandler(ctx, repoTransaction)
	handlerTransactionDetail := transaction_handler.NewTransactionDetailHandler(ctx, repoTransactionDetail)
	r.HandleFunc("/store-customer", handlerCustomer.StoreController).Methods(http.MethodPost)
	r.HandleFunc("/store-transaction", handlerTransaction.GetAllTransaction).Methods(http.MethodGet)
	r.HandleFunc("/store-transaction-detail", handlerTransactionDetail.GetAllTransactionDetail).Methods(http.MethodGet)

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)
}
