package transaction_test

import (
	"fmt"
	"game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/domain/entity/transaction"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
Positif Case
*/
func TestNewTransaction(t *testing.T) {

	itemList, _ := item.NewListItem([]item.DTOItem{
		{
			Id:       1,
			Nama:     "Xbox",
			Kategori: "Service",
			Harga:    350000,
			Jumlah:   1,
		},
		{
			Id:       6,
			Nama:     "Xbox",
			Kategori: "Buy Console",
			Harga:    5432100,
			Jumlah:   1,
		},
	})

	transaction, err := transaction.NewTransaction(transaction.DTOTransaction{
		Id:                1,
		VoucherCustomerId: 1,
		CustomerId:        25123123,
		CodeTransaction:   time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian:  time.Now().String(),
		Total:             300000,
		HargaDiscount:     100000,
		TotalHarga:        200000,
		ItemPembelian:     itemList,
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
		CodeTransaction:   time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian:  time.Now().String(),
		Total:             300000,
		HargaDiscount:     100000,
		TotalHarga:        200000,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "ID COSTOMER NOT SET", err.Error())
}
