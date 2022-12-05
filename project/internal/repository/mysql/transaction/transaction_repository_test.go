package transaction_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
