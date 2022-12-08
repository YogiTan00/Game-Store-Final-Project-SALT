package transaction_detail_test

import (
	"context"
	"fmt"
	"game-store-final-project/project/internal/repository/mysql/transaction_detail"
	"game-store-final-project/project/pkg/mysql_connection"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mysqlConn                  = mysql_connection.InitMysqlDB()
	ctx                        = context.Background()
	repoMysqlTransactionDetail = transaction_detail.NewTransactionDetailMysqlInteractor(mysqlConn)
)

func TestTransactionDetailRepositoryMysqlInteractor_GetTransaction(t *testing.T) {
	transactionDetail, err := repoMysqlTransactionDetail.GetAllTransactionDetail(ctx)
	fmt.Println(transactionDetail)
	assert.NotNil(t, transactionDetail)
	assert.Nil(t, err)
}

func TestTransactionDetailRepositoryMysqlInteractor_GetTransactionDetailByID(t *testing.T) {
	transactionDetail, err := repoMysqlTransactionDetail.GetTransactionDetailByID(ctx, "1")
	fmt.Println(transactionDetail)
	assert.NotNil(t, transactionDetail)
	assert.Nil(t, err)
}

func TestTransactionDetailRepositoryMysqlInteractor_GetAllTransactionDetailByID(t *testing.T) {
	transactionDetail, err := repoMysqlTransactionDetail.GetAllTransactionDetailByID(ctx, "1")
	fmt.Println(transactionDetail)
	assert.NotNil(t, transactionDetail)
	assert.Nil(t, err)
}
