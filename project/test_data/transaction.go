package test_data

import (
	item2 "game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/domain/entity/transaction_detail"
	"time"
)

func GetTestDataTransaction() *transaction.Transaction {
	item, _ := item2.NewItem(item2.DTOItem{
		Id:       1,
		Nama:     "Xbox",
		Kategori: "Service",
		Harga:    350000,
		Jumlah:   1,
	})

	listTransDetail, _ := transaction_detail.NewListTransactionDetail([]transaction_detail.DTOTransactionDetail{{
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
	transaction, _ := transaction.NewTransaction(transaction.DTOTransaction{
		Id:               1,
		CustomerId:       25123123,
		CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian: &tglPembelian,
		Total:            300000,
		TransDetail:      listTransDetail,
	})
	return transaction
}

func GetTestDataCountTransaction(count int) []*transaction.Transaction {
	listTrans := make([]*transaction.Transaction, 0)
	for i := 1; i <= count; i++ {
		item, _ := item2.NewItem(item2.DTOItem{
			Id:       1,
			Nama:     "Xbox",
			Kategori: "Service",
			Harga:    350000,
			Jumlah:   1,
		})

		listTransDetail, _ := transaction_detail.NewListTransactionDetail([]transaction_detail.DTOTransactionDetail{{
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
		transaction, _ := transaction.NewTransaction(transaction.DTOTransaction{
			Id:               1,
			CustomerId:       25123123,
			CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
			Tanggalpembelian: &tglPembelian,
			Total:            300000,
			TransDetail:      listTransDetail,
		})
		listTrans = append(listTrans, transaction)
	}
	return listTrans
}
