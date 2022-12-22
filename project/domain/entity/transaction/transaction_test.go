package transaction_test

import (
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
	listTransD := make([]*transaction_detail2.TransactionDetail, 3)
	for range listTransD {
		item, errItem := item.NewItem(item.DTOItem{
			Id:       1,
			Nama:     "Xbox",
			Kategori: "Service Console",
			Harga:    350000,
			Jumlah:   1,
		})

		transD, errTransD := transaction_detail2.NewTransactionDetail(transaction_detail2.DTOTransactionDetail{
			Id:              1,
			TransactionId:   1,
			ItemId:          1,
			DetailItem:      item,
			JumlahPembelian: 1,
			HargaPembelian:  350000,
			HargaDiscount:   0,
			Total:           350000,
		})
		listTransD = append(listTransD, transD)
		assert.Nil(t, errItem)
		assert.Nil(t, errTransD)
	}

	tglPembelian := time.Now()
	transaction, err := transaction.NewTransaction(transaction.DTOTransaction{
		Id:               1,
		CustomerId:       25123123,
		CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian: &tglPembelian,
		Total:            350000,
		TransDetail:      listTransD,
	})

	assert.Nil(t, err)
	assert.Equal(t, 25123123, transaction.GetCustomerID())
}

func TestNewTransactionWithDetail(t *testing.T) {
	listTransD := make([]*transaction_detail2.TransactionDetail, 3)
	for range listTransD {
		item, errItem := item.NewItem(item.DTOItem{
			Id:       1,
			Nama:     "Xbox",
			Kategori: "Service Console",
			Harga:    350000,
			Jumlah:   1,
		})

		transD, errTransD := transaction_detail2.NewTransactionDetail(transaction_detail2.DTOTransactionDetail{
			Id:              1,
			TransactionId:   1,
			ItemId:          1,
			DetailItem:      item,
			JumlahPembelian: 1,
			HargaPembelian:  350000,
			HargaDiscount:   0,
			Total:           350000,
		})
		listTransD = append(listTransD, transD)
		assert.Nil(t, errItem)
		assert.Nil(t, errTransD)
	}

	tglPembelian := time.Now()
	transaction, err := transaction.NewTransactionWithDetail(transaction.DTOTransaction{
		Id:               1,
		CustomerId:       25123123,
		CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian: &tglPembelian,
		Total:            350000,
		TransDetail:      listTransD,
	}, listTransD)

	assert.Nil(t, err)
	assert.Equal(t, 25123123, transaction.GetCustomerID())
}

/*
Negative Case
*/

func TestValidateTransactionWithDetailNoHaveID(t *testing.T) {
	listTransD := make([]*transaction_detail2.TransactionDetail, 3)
	for range listTransD {
		item, errItem := item.NewItem(item.DTOItem{
			Id:       1,
			Nama:     "Xbox",
			Kategori: "Service Console",
			Harga:    350000,
			Jumlah:   1,
		})

		transD, errTransD := transaction_detail2.NewTransactionDetail(transaction_detail2.DTOTransactionDetail{
			Id:              1,
			TransactionId:   1,
			ItemId:          1,
			DetailItem:      item,
			JumlahPembelian: 1,
			HargaPembelian:  350000,
			HargaDiscount:   0,
			Total:           350000,
		})
		listTransD = append(listTransD, transD)
		assert.Nil(t, errItem)
		assert.Nil(t, errTransD)
	}
	tglPembelian := time.Now()
	_, err := transaction.NewTransaction(transaction.DTOTransaction{
		Id:               1,
		CustomerId:       0,
		CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian: &tglPembelian,
		Total:            300000,
		TransDetail:      listTransD,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "ID COSTOMER NOT SET", err.Error())
}

func TestValidateTransactionID(t *testing.T) {
	listTransD := make([]*transaction_detail2.TransactionDetail, 3)
	for range listTransD {
		item, errItem := item.NewItem(item.DTOItem{
			Id:       1,
			Nama:     "Xbox",
			Kategori: "Service Console",
			Harga:    350000,
			Jumlah:   1,
		})

		transD, errTransD := transaction_detail2.NewTransactionDetail(transaction_detail2.DTOTransactionDetail{
			Id:              1,
			TransactionId:   1,
			ItemId:          1,
			DetailItem:      item,
			JumlahPembelian: 1,
			HargaPembelian:  350000,
			HargaDiscount:   0,
			Total:           350000,
		})
		listTransD = append(listTransD, transD)
		assert.Nil(t, errItem)
		assert.Nil(t, errTransD)
	}
	tglPembelian := time.Now()
	_, err := transaction.NewTransactionWithDetail(transaction.DTOTransaction{
		Id:               1,
		CustomerId:       0,
		CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian: &tglPembelian,
		Total:            300000,
		TransDetail:      listTransD,
	}, listTransD)

	assert.NotNil(t, err)
	assert.Equal(t, "ID COSTOMER NOT SET", err.Error())
}
