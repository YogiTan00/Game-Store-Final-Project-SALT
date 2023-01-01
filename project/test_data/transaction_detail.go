package test_data

import (
	"game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/domain/entity/transaction_detail"
)

func GetTestDataTransactionDetail() *transaction_detail.TransactionDetail {
	item, _ := item.NewItem(item.DTOItem{
		Id:       1,
		Nama:     "Xbox",
		Kategori: "Service",
		Harga:    350000,
		Jumlah:   1,
	})
	transaction, _ := transaction_detail.NewTransactionDetail(transaction_detail.DTOTransactionDetail{
		Id:              1,
		TransactionId:   1,
		ItemId:          1,
		DetailItem:      item,
		JumlahPembelian: 1,
		HargaPembelian:  4360000,
		HargaDiscount:   0,
		Total:           4360000,
	})

	return transaction
}

func GetTestDataCountTransactionDetail(countTransD int) []*transaction_detail.TransactionDetail {
	listTransactionDetail := make([]*transaction_detail.TransactionDetail, 0)
	for i := 1; i <= countTransD; i++ {
		item, _ := item.NewItem(item.DTOItem{
			Id:       1,
			Nama:     "Xbox",
			Kategori: "Service",
			Harga:    350000,
			Jumlah:   1,
		})
		transaction, _ := transaction_detail.NewTransactionDetail(transaction_detail.DTOTransactionDetail{
			Id:              1,
			TransactionId:   1,
			ItemId:          1,
			DetailItem:      item,
			JumlahPembelian: 1,
			HargaPembelian:  4360000,
			HargaDiscount:   0,
			Total:           4360000,
		})
		listTransactionDetail = append(listTransactionDetail, transaction)
	}
	return listTransactionDetail
}
