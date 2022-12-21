package test_data

import (
	item2 "game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/domain/entity/transaction_detail"
	"time"
)

func GetTestDataTransaction(countTransD int) *transaction.Transaction {

	listTransD := make([]*transaction_detail.TransactionDetail, 0)
	for i := 0; i <= countTransD; i++ {
		item, _ := item2.NewItem(item2.DTOItem{
			Id:       1,
			Nama:     "Xbox",
			Kategori: "Service",
			Harga:    350000,
			Jumlah:   1,
		})

		TransDetail, _ := transaction_detail.NewTransactionDetail(transaction_detail.DTOTransactionDetail{
			Id:              1,
			TransactionId:   1,
			ItemId:          1,
			DetailItem:      item,
			JumlahPembelian: 1,
			HargaPembelian:  350000,
			HargaDiscount:   0,
			Total:           350000,
		})
		listTransD = append(listTransD, TransDetail)
	}

	tglPembelian := time.Now()
	transaction, _ := transaction.NewTransaction(transaction.DTOTransaction{
		Id:               1,
		CustomerId:       25123123,
		CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian: &tglPembelian,
		Total:            300000,
		TransDetail:      listTransD,
	})
	return transaction
}

func GetTestDataCountTransaction(countTrans int, countTransD int) []*transaction.Transaction {
	listTrans := make([]*transaction.Transaction, 0)
	for i := 1; i <= countTrans; i++ {
		listTransD := make([]*transaction_detail.TransactionDetail, 0)
		for j := 0; j <= countTransD; j++ {
			item, _ := item2.NewItem(item2.DTOItem{
				Id:       1,
				Nama:     "Xbox",
				Kategori: "Service",
				Harga:    350000,
				Jumlah:   1,
			})

			TransDetail, _ := transaction_detail.NewTransactionDetail(transaction_detail.DTOTransactionDetail{
				Id:              1,
				TransactionId:   1,
				ItemId:          1,
				DetailItem:      item,
				JumlahPembelian: 1,
				HargaPembelian:  350000,
				HargaDiscount:   0,
				Total:           350000,
			})
			listTransD = append(listTransD, TransDetail)
		}

		tglPembelian := time.Now()
		transaction, _ := transaction.NewTransaction(transaction.DTOTransaction{
			Id:               1,
			CustomerId:       25123123,
			CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
			Tanggalpembelian: &tglPembelian,
			Total:            300000,
			TransDetail:      listTransD,
		})
		listTrans = append(listTrans, transaction)
	}
	return listTrans
}
