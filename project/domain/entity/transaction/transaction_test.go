package transaction_test

import (
	"fmt"
	"game-store-final-project/project/domain/entity/transaction"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Positif Case
*/
func TestNewTransaction(t *testing.T) {
	transaction, err := transaction.NewTransaction(transaction.DTOTransaction{
		Id:                1,
		VoucherCustomerId: 1,
		CustomerId:        25123123,
		CodeTransaction:   "INV/20221206DPFddswJ",
		Tanggalpembelian:  "2022-10-10",
		Total:             300000,
		HargaDiscount:     100000,
		TotalHarga:        200000,
	})
	fmt.Println(transaction)
	assert.Nil(t, err)
	assert.Equal(t, 25123123, transaction.GetCustomerID())
}

/*
Negative Case
*/

func TestValidateTransactionID(t *testing.T) {
	_, err := transaction.NewTransaction(transaction.DTOTransaction{
		Id:                1,
		VoucherCustomerId: 1,
		CustomerId:        0,
		CodeTransaction:   "INV/20221206DPFddswJ",
		Tanggalpembelian:  "2022-10-10",
		Total:             300000,
		HargaDiscount:     100000,
		TotalHarga:        200000,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "ID COSTOMER NOT SET", err.Error())
}
