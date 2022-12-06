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
		Id:              1,
		CodeTransaction: "INV05122022190500",
		TransactionId:   1,
		ItemId:          "1",
		JumlahPembelian: 1,
		HargaPembelian:  4360000,
		Total:           4360000,
	})

	assert.Nil(t, err)
	assert.Equal(t, 1, transaction.GetTransactionID())
}

/*
Negative Case
*/

func TestValidateTransactionDetails_CodeTransaction(t *testing.T) {
	_, err := transaction.NewTransactionDetail(transaction.DTOTransactionDetail{
		Id:              1,
		CodeTransaction: "",
		TransactionId:   1,
		ItemId:          "1",
		JumlahPembelian: 1,
		HargaPembelian:  4360000,
		Total:           4360000,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "CODE TRANSACTION NOT SET", err.Error())
}

func TestValidateTransactionDetails_TransactionID(t *testing.T) {
	_, err := transaction.NewTransactionDetail(transaction.DTOTransactionDetail{
		Id:              1,
		CodeTransaction: "INV05122022190500",
		TransactionId:   0,
		ItemId:          "1",
		JumlahPembelian: 1,
		HargaPembelian:  4360000,
		Total:           4360000,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "TRANSACTION ID NOT SET", err.Error())
}

func TestValidateTransactionDetails_ItemID(t *testing.T) {
	_, err := transaction.NewTransactionDetail(transaction.DTOTransactionDetail{
		Id:              1,
		CodeTransaction: "INV05122022190500",
		TransactionId:   1,
		ItemId:          "",
		JumlahPembelian: 1,
		HargaPembelian:  4360000,
		Total:           4360000,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "ITEM ID NOT SET", err.Error())
}
