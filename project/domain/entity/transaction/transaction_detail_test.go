package transaction_test

import (
	"game-store-final-project/project/domain/entity/transaction"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Positif Case
*/
func TestNewTransactionDetail(t *testing.T) {
	transaction, err := transaction.NewTransactionDetail(transaction.DTOTransactionDetail{
		Id:            1,
		TransactionId: 1,
		ItemId:        6,
	})

	assert.Nil(t, err)
	assert.Equal(t, 1, transaction.GetTransactionID())
}

/*
Negative Case
*/

func TestValidateTransactionDetails_TransactionID(t *testing.T) {
	_, err := transaction.NewTransactionDetail(transaction.DTOTransactionDetail{
		Id:            1,
		TransactionId: 0,
		ItemId:        6,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "ID TRANSACTION NOT SET", err.Error())
}

func TestValidateTransactionDetails_ItemID(t *testing.T) {
	_, err := transaction.NewTransactionDetail(transaction.DTOTransactionDetail{
		Id:            1,
		TransactionId: 1,
		ItemId:        0,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "ID ITEM NOT SET", err.Error())
}
