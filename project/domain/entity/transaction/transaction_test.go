package transaction_test

import (
	"fmt"
	"game-store-final-project/project/domain/entity/transaction"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
Positif Case
*/
func TestNewTransaction(t *testing.T) {
	tglPembelian := time.Now()
	transaction, err := transaction.NewTransaction(transaction.DTOTransaction{
		Id:               1,
		CustomerId:       25123123,
		CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian: &tglPembelian,
		Total:            300000,
	})
	fmt.Println(transaction)
	assert.Nil(t, err)
	assert.Equal(t, 25123123, transaction.GetCustomerID())
}

/*
Negative Case
*/

func TestValidateTransactionID(t *testing.T) {
	tglPembelian := time.Now()
	_, err := transaction.NewTransaction(transaction.DTOTransaction{
		Id:               1,
		CustomerId:       0,
		CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian: &tglPembelian,
		Total:            300000,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "ID COSTOMER NOT SET", err.Error())
}
