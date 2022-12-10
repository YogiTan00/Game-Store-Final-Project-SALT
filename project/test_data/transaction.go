package test_data

import (
	"game-store-final-project/project/domain/entity/transaction"
	"time"
)

func GetTestDataTransaction() *transaction.Transaction {
	tglPembelian := time.Now()
	transaction, _ := transaction.NewTransaction(transaction.DTOTransaction{
		Id:               1,
		CustomerId:       25123123,
		CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
		Tanggalpembelian: &tglPembelian,
		Total:            300000,
	})
	return transaction
}

func GetTestDataCountTransaction(count int) []*transaction.Transaction {
	listTrans := make([]*transaction.Transaction, 0)
	for i := 1; i <= count; i++ {
		tglPembelian := time.Now()
		transaction, _ := transaction.NewTransaction(transaction.DTOTransaction{
			Id:               1,
			CustomerId:       25123123,
			CodeTransaction:  time.Now().Format("INV02D01M2006Y15H04M05S"),
			Tanggalpembelian: &tglPembelian,
			Total:            300000,
		})
		listTrans = append(listTrans, transaction)
	}
	return listTrans
}
