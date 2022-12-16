package transaction_test

import (
	"fmt"
	"game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/domain/entity/transaction"
	transaction_detail2 "game-store-final-project/project/domain/entity/transaction_detail"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
Positif Case
*/
func TestNewTransaction(t *testing.T) {
	item, _ := item.NewItem(item.DTOItem{
		Id:       0,
		Nama:     "",
		Kategori: "",
		Harga:    0,
		Jumlah:   0,
	})

	listTransDetail, _ := transaction_detail2.NewListTransactionDetail([]transaction_detail2.DTOTransactionDetail{{
		Id:              1,
		TransactionId:   1,
		ItemId:          1,
		DetailItem:      item,
		JumlahPembelian: 1,
		HargaPembelian:  350000,
		HargaDiscount:   0,
		Total:           350000,
	}})

	tglPembelian := time.Now()
	transaction, err := transaction.NewTransaction(transaction.DTOTransaction{
		Id:               1,
		CustomerId:       25123123,
		CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian: &tglPembelian,
		TransDetail:      listTransDetail,
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
