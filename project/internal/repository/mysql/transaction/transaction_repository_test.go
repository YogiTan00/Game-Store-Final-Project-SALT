package transaction_test

import (
	"context"
	"fmt"
	"game-store-final-project/project/internal/repository/mysql/transaction"
	transaction_detail2 "game-store-final-project/project/internal/repository/mysql/transaction_detail"
	"game-store-final-project/project/pkg/mysql_connection"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

var (
	mysqlConn            = mysql_connection.InitMysqlDB()
	ctx                  = context.Background()
	repoMysqlTransaction = transaction.NewTransactionMysqlInteractor(mysqlConn)
	repoMysqlTransDetail = transaction_detail2.NewTransactionDetailMysqlInteractor(mysqlConn)
)

func TestTransactionRepositoryMysqlInteractor_GetTransactionByID(t *testing.T) {
	transaction, err := repoMysqlTransaction.GetTransactionByID(ctx, "3")
	listTransD, errTransD := repoMysqlTransDetail.GetAllTransactionDetailByID(ctx, strconv.Itoa(transaction.GetID()))
	fmt.Println(transaction)
	fmt.Println(listTransD)
	assert.NotNil(t, transaction)
	assert.NotNil(t, listTransD)
	assert.Nil(t, err)
	assert.Nil(t, errTransD)
}

func TestTransactionRepositoryMysqlInteractor_GetAllTransaction(t *testing.T) {
	transaction, err := repoMysqlTransaction.GetAllTransaction(ctx)
	for _, data := range transaction {
		transD, errTransD := repoMysqlTransDetail.GetAllTransactionDetailByID(ctx, strconv.Itoa(data.GetID()))
		data.AddTransDetail(transD)
		assert.NotNil(t, transD)
		assert.Nil(t, errTransD)
		fmt.Println("TransD : ", transD)
	}
	fmt.Println(transaction)
	assert.NotNil(t, transaction)
	assert.Nil(t, err)
}

func TestTransactionRepositoryMysqlInteractor_GetAllTransactionByCustomerID(t *testing.T) {
	transaction, err := repoMysqlTransaction.GetAllTransactionByCustomerID(ctx, "1")
	for _, data := range transaction {
		transD, errTransD := repoMysqlTransDetail.GetAllTransactionDetailByID(ctx, strconv.Itoa(data.GetID()))
		data.AddTransDetail(transD)
		assert.NotNil(t, transD)
		assert.Nil(t, errTransD)
		fmt.Println("TransD : ", transD)
	}
	fmt.Println(transaction)
	assert.NotNil(t, transaction)
	assert.Nil(t, err)
}
