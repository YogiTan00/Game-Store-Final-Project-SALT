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
