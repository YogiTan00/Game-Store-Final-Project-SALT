package transaction_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTransactionRepositoryMysqlInteractor_GetTransaction(t *testing.T) {
	transaction, err := repoMysqlTransaction.GetAllTransaction(ctx)
	fmt.Println(transaction)
	assert.NotNil(t, transaction)
	assert.Nil(t, err)
}
