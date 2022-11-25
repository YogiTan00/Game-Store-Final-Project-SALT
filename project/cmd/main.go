package main

import (
	"context"
	"fmt"
	"game-store-final-project/project/internal/delivery/http/customer_handler"
	repo "game-store-final-project/project/internal/repository/mysql"
	"game-store-final-project/project/pkg/mysql_connection"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	mysqlConn    = mysql_connection.InitMysqlDB()
	repoCustomer = repo.NewArticleRepositoryMysqlInteractor(mysqlConn)
	ctx          = context.Background()
)

func main() {
	r := mux.NewRouter()

	// routes student
	handlerCustomer := customer_handler.NewCustomerHandler(ctx, repoCustomer)
	r.HandleFunc("/store-customer", handlerCustomer.Store).Methods(http.MethodPost)

	fmt.Println("localhost:8080")
	http.ListenAndServe(":8080", r)
}
