package transaction_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
