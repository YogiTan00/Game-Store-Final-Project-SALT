package transaction_detail_test

import (
	"game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/domain/entity/transaction_detail"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Positif Case
*/
func TestNewTransactionDetail(t *testing.T) {

	item, _ := item.NewItem(item.DTOItem{
		Id:       1,
		Nama:     "Xbox",
		Kategori: "Service",
		Harga:    350000,
		Jumlah:   1,
	})
	transaction, err := transaction_detail.NewTransactionDetail(transaction_detail.DTOTransactionDetail{
		Id:              1,
		TransactionId:   1,
		ItemId:          "1",
		DetailItem:      item,
		JumlahPembelian: 1,
		HargaPembelian:  4360000,
		HargaDiscount:   0,
		Total:           4360000,
	})

	assert.Nil(t, err)
	assert.Equal(t, 1, transaction.GetTransactionID())
}

/*
Negative Case
*/

func TestValidateTransactionDetails_TransactionID(t *testing.T) {
	item, _ := item.NewItem(item.DTOItem{
		Id:       1,
		Nama:     "Xbox",
		Kategori: "Service",
		Harga:    350000,
		Jumlah:   1,
	})
	_, err := transaction_detail.NewTransactionDetail(transaction_detail.DTOTransactionDetail{
		Id:              1,
		TransactionId:   0,
		ItemId:          "1",
		DetailItem:      item,
		JumlahPembelian: 1,
		HargaPembelian:  4360000,
		HargaDiscount:   0,
		Total:           4360000,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "TRANSACTION ID NOT SET", err.Error())
}

func TestValidateTransactionDetails_ItemID(t *testing.T) {
	item, _ := item.NewItem(item.DTOItem{
		Id:       1,
		Nama:     "Xbox",
		Kategori: "Service",
		Harga:    350000,
		Jumlah:   1,
	})
	_, err := transaction_detail.NewTransactionDetail(transaction_detail.DTOTransactionDetail{
		Id:              1,
		TransactionId:   1,
		ItemId:          "",
		DetailItem:      item,
		JumlahPembelian: 1,
		HargaPembelian:  4360000,
		HargaDiscount:   0,
		Total:           4360000,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "ITEM ID NOT SET", err.Error())
}
