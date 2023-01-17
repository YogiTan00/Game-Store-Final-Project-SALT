package main

import (
	"context"
	"fmt"
	"game-store-final-project/project/internal/config/database/mysql"
	"game-store-final-project/project/internal/delivery/http/customer_handler"
	"game-store-final-project/project/internal/delivery/http/item_handler"
	"game-store-final-project/project/internal/delivery/http/transaction_handler"
	"game-store-final-project/project/internal/delivery/http/voucher_handler"
	repo "game-store-final-project/project/internal/repository/mysql/customer"
	item2 "game-store-final-project/project/internal/repository/mysql/item"
	"game-store-final-project/project/internal/repository/mysql/transaction"
	transaction_detail2 "game-store-final-project/project/internal/repository/mysql/transaction_detail"
	repo_voucher "game-store-final-project/project/internal/repository/mysql/voucher"
	usecase_cust "game-store-final-project/project/internal/usecase/customer"
	"game-store-final-project/project/internal/usecase/item"
	usecase_trx "game-store-final-project/project/internal/usecase/transaction"
	"game-store-final-project/project/internal/usecase/transaction_detail"
	"game-store-final-project/project/internal/usecase/voucher"
	"game-store-final-project/project/internal/usecase/voucher_customer"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	ctx                      = context.Background()
	mysqlConn                = mysql.InitMysqlDB()
	repoCustomer             = repo.NewCustomerRepositoryMysqlInteractor(mysqlConn)
	repoTransaction          = transaction.NewTransactionMysqlInteractor(mysqlConn)
	repoTransactionDetail    = transaction_detail2.NewTransactionDetailMysqlInteractor(mysqlConn)
	repoItem                 = item2.NewItemMysqlInteractor(mysqlConn)
	repoVoucher              = repo_voucher.NewVoucherRepositoryMysqlInteractor(mysqlConn)
	useCaseCustomer          = usecase_cust.NewCustomerUseCaseInteractor(ctx, repoCustomer)
	useCaseTransaction       = usecase_trx.NewTransactionUseCaseInteractor(ctx, repoTransaction, repoItem, repoVoucher, repoTransactionDetail, repoCustomer)
	useCaseTransactionDetail = transaction_detail.NewTransactionDetailUseCaseInteractor(repoTransactionDetail, repoItem)
	useCaseItem              = item.NewItemUseCaseInteractor(repoItem)
	useCaseVoucher           = voucher.NewVoucherUseCaseInteractor(repoVoucher)
	useCaseVoucherCustomer   = voucher_customer.NewVoucherCustomerUseCaseInteractor(repoVoucher, repoCustomer)
)

func main() {
	r := mux.NewRouter()

	// routes customer
	handlerCustomer := customer_handler.NewUseCaseCustomerHandler(useCaseCustomer)
	handlerTrx := transaction_handler.NewUsecaseTransactionHandler(useCaseTransaction)
	handlerTransaction := transaction_handler.NewUseCaseTransactionHandler(useCaseTransaction, useCaseTransactionDetail, useCaseItem, useCaseVoucher)
	handlerTransactionDetail := transaction_handler.NewUseCaseTransactionDetailHandler(useCaseTransactionDetail, useCaseItem)
	handlerItem := item_handler.NewuseCaseItemHandler(useCaseItem)
	handlerVoucher := voucher_handler.NewVoucherHandler(useCaseVoucherCustomer, useCaseCustomer)
	// customer
	r.HandleFunc("/customer/list-trx/{nik}", handlerCustomer.IndexController).Methods(http.MethodGet)
	r.HandleFunc("/customer/store", handlerCustomer.StoreController).Methods(http.MethodPost)

	// transaksi
	r.HandleFunc("/transaction/store", handlerTrx.StoreController).Methods(http.MethodPost)
	r.HandleFunc("/get-transaction", handlerTransaction.GetAllTransactionHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-transaction/{id}", handlerTransaction.GetTransactionByIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-transaction/customer/{id}", handlerTransaction.GetAllTransactionByCustomerIDHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-transaction-detail", handlerTransactionDetail.GetAllTransactionDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-transaction-detail/{id}", handlerTransactionDetail.GetAllTransactionDetailByIDHandler).Methods(http.MethodGet)

	// item
	r.HandleFunc("/get-item", handlerItem.GetAllItemHandler).Methods(http.MethodGet)
	r.HandleFunc("/get-item/{id}", handlerItem.GetItemByIDHandler).Methods(http.MethodGet)

	// voucher
	r.HandleFunc("/get-voucher/{id}", handlerVoucher.GetVoucherCustomerByIdHandler).Methods(http.MethodGet)

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)
}
