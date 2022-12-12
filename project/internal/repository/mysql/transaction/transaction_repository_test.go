package transaction_test

import (
	"context"
	"fmt"
	"game-store-final-project/project/internal/repository/mysql/transaction"
	"game-store-final-project/project/pkg/mysql_connection"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mysqlConn            = mysql_connection.InitMysqlDB()
	ctx                  = context.Background()
	repoMysqlTransaction = transaction.NewTransactionMysqlInteractor(mysqlConn)
)

func TestTransactionRepositoryMysqlInteractor_GetTransactionByID(t *testing.T) {
	transaction, err := repoMysqlTransaction.GetTransactionByID(ctx, "1")
	fmt.Println(transaction)
	assert.NotNil(t, transaction)
	assert.Nil(t, err)
}

func TestTransactionRepositoryMysqlInteractor_GetAllTransaction(t *testing.T) {
	transaction, err := repoMysqlTransaction.GetAllTransaction(ctx)
	fmt.Println(transaction)
	assert.NotNil(t, transaction)
	assert.Nil(t, err)
}

func TestTransactionRepositoryMysqlInteractor_GetAllTransactionByID(t *testing.T) {
	transaction, err := repoMysqlTransaction.GetAllTransactionByID(ctx, "1")
	fmt.Println(transaction)
	assert.NotNil(t, transaction)
	assert.Nil(t, err)
}
