package main

import (
	"context"
	"fmt"
	"game-store-final-project/project/internal/delivery/http/customer_handler"
	"game-store-final-project/project/internal/delivery/http/item_handler"
	"game-store-final-project/project/internal/delivery/http/transaction_handler"
	repo "game-store-final-project/project/internal/repository/mysql/customer"
	item2 "game-store-final-project/project/internal/repository/mysql/item"
	"game-store-final-project/project/internal/repository/mysql/transaction"
	transaction_detail2 "game-store-final-project/project/internal/repository/mysql/transaction_detail"
	repo_voucher "game-store-final-project/project/internal/repository/mysql/voucher"
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
	repoTransactionDetail = transaction_detail2.NewTransactionDetailMysqlInteractor(mysqlConn)
	repoItem              = item2.NewItemMysqlInteractor(mysqlConn)
	repoVoucher           = repo_voucher.NewVoucherRepositoryMysqlInteractor(mysqlConn)
	useCaseCustomer       = usecase_cust.NewCustomerUseCaseInteractor(ctx, repoCustomer)
	useCaseTransaction    = usecase_trx.NewTransactionUseCaseInteractor(ctx, repoTransaction, repoItem, repoVoucher, repoTransactionDetail)
	//useCaseTransactionDetail = transaction_detail.NewTransactionDetailUseCaseInteractor(repoTransactionDetail)
	//useCaseItem = item.NewItemUseCaseInteractor(repoItem)
)

func main() {
	r := mux.NewRouter()

	// routes customer
	handlerCustomer := customer_handler.NewCustomerHandler(useCaseCustomer)
	handlerTrx := transaction_handler.NewUsecaseTransactionHandler(useCaseTransaction)
	handlerTransaction := transaction_handler.NewTransactionHandler(repoTransaction, repoTransactionDetail, repoItem, repoVoucher)
	handlerTransactionDetail := transaction_handler.NewTransactionDetailHandler(repoTransactionDetail, repoItem)
	handlerItem := item_handler.NewItemHandler(repoItem)
	// customer
	r.HandleFunc("/customer/list-trx/{nik}", handlerCustomer.IndexController).Methods(http.MethodGet)
	r.HandleFunc("/customer/store", handlerCustomer.StoreController).Methods(http.MethodPost)

	// transaksi
	r.HandleFunc("/transaction/store", handlerTrx.StoreController).Methods(http.MethodPost)
	r.HandleFunc("/get-transaction", handlerTransaction.GetAllTransactionHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-transaction/{id}", handlerTransaction.GetTransactionByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-transaction/customer/{id}", handlerTransaction.GetAllTransactionByCustomerIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-transaction-detail", handlerTransactionDetail.GetAllTransactionDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-transaction-detail/{id}", handlerTransactionDetail.GeAllTransactionDetailByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-item", handlerItem.GetAllItemHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-item/{id}", handlerItem.GetItemByIDHandler).Methods(http.MethodGet)

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)
}
