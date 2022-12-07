package main

import (
	"context"
	"fmt"
	"game-store-final-project/project/internal/delivery/http/customer_handler"
	"game-store-final-project/project/internal/delivery/http/item"
	"game-store-final-project/project/internal/delivery/http/transaction_handler"
	repo "game-store-final-project/project/internal/repository/mysql/customer"
	item2 "game-store-final-project/project/internal/repository/mysql/item"
	"game-store-final-project/project/internal/repository/mysql/transaction"
	usecase_cust "game-store-final-project/project/internal/usecase/customer"
	usecase_trx "game-store-final-project/project/internal/usecase/transaction"
	"game-store-final-project/project/pkg/mysql_connection"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ctx                   = context.Background()
	mysqlConn             = mysql_connection.InitMysqlDB()
	repoCustomer          = repo.NewCustomerRepositoryMysqlInteractor(mysqlConn)
	repoTransaction       = transaction.NewTransactionMysqlInteractor(mysqlConn)
	repoItem              = item2.NewItemMysqlInteractor(mysqlConn)
	useCaseCustomer       = usecase_cust.NewCustomerUseCaseInteractor(ctx, repoCustomer)
	useCaseTransaction    = usecase_trx.NewTransactionUseCaseInteractor(ctx, repoTransaction, repoItem)
	repoTransactionDetail = transaction.NewTransactionDetailMysqlInteractor(mysqlConn)
)

func main() {
	r := mux.NewRouter()

	// routes customer
	handlerCustomer := customer_handler.NewCustomerHandler(useCaseCustomer)
	handlerTrx := transaction_handler.NewUsecaseTransactionHandler(useCaseTransaction)
	handlerTransaction := transaction_handler.NewTransactionHandler(ctx, repoTransaction)

	handlerTransactionDetail := transaction_handler.NewTransactionDetailHandler(ctx, repoTransactionDetail, repoItem)
	handlerItem := item.NewItemHandler(repoItem)
	// customer
	r.HandleFunc("/store-customer", handlerCustomer.StoreController).Methods(http.MethodPost)

	// transaksi
	r.HandleFunc("/transaction/store", handlerTrx.StoreController).Methods(http.MethodPost)
	r.HandleFunc("/get-transaction", handlerTransaction.GetAllTransaction).Methods(http.MethodGet)
	r.HandleFunc("/get-transaction/{id}", handlerTransaction.GetAllTransactionByID).Methods(http.MethodGet)
	r.HandleFunc("/get-transaction-detail", handlerTransactionDetail.GetAllTransactionDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-transaction-detail/{id}", handlerTransactionDetail.GeAllTransactionDetailByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-item", handlerItem.GetAllItem).Methods(http.MethodGet)
	r.HandleFunc("/get-item/{id}", handlerItem.GetAllItemByID).Methods(http.MethodGet)

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)
}
