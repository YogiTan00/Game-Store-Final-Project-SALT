package test_data

import (
	"game-store-final-project/project/domain/entity/customer"
	item2 "game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/domain/entity/transaction_detail"
	"time"
)

func GetTestDataCustomer() *customer.Customer {
	customer, _ := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "3204223423442582",
		Nama:         "Taupik Pirdian",
		Alamat:       "Bandung",
		NoTlp:        "085846342122",
		JenisKelamin: "LK",
	})

	return customer
}

func GetTestDataCountCustomer(count int) []*customer.Customer {
	listCustomer := make([]*customer.Customer, 0)
	for i := 1; i <= count; i++ {
		customer, _ := customer.NewCustomer(customer.DTOCustomer{
			Nik:          "3204223423442582",
			Nama:         "Taupik Pirdian",
			Alamat:       "Bandung",
			NoTlp:        "085846342122",
			JenisKelamin: "LK",
		})
		listCustomer = append(listCustomer, customer)
	}
	return listCustomer
}

func DataCustomerWithTransaction() *customer.Customer {
	listTransD := make([]*transaction_detail.TransactionDetail, 0)
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

	tglPembelian := time.Now()
	trxSlice := make([]*transaction.Transaction, 0)
	transaction, _ := transaction.NewTransaction(transaction.DTOTransaction{
		Id:               1,
		CustomerId:       25123123,
		CodeTransaction:  "INV-20D12M2022Y14H16M17S",
		Tanggalpembelian: &tglPembelian,
		Total:            350000,
		TransDetail:      listTransD,
	})
	trxSlice = append(trxSlice, transaction)

	customer, _ := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "3204223423442582",
		Nama:         "Taupik Pirdian",
		Alamat:       "Bandung",
		NoTlp:        "085846342122",
		JenisKelamin: "LK",
	})

	customer.AddTrx(trxSlice)

	return customer
}
