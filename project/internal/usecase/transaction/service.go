package transaction

import (
	"fmt"
	"game-store-final-project/project/domain/entity/transaction"
	"time"
)

func (cu *TransactionUseCaseInteractor) StoreTransaction(customer_id int, tanggal_pembelian string, voucher string, items []*transaction.DTODetailTransaction) (*transaction.Transaction, error) {
	// Debug
	// for _, item := range items {
	// 	fmt.Println(*item)
	// }

	// convert string to time
	date, _ := time.Parse("2006-01-02", tanggal_pembelian)
	// build data to entity
	transaction, err := transaction.NewBuildTransaction(customer_id, date, voucher)
	if err != nil {
		return nil, err
	}

	fmt.Println(transaction)

	// errInsert := cu.RepoCustomer.StoreCustomer(cu.ctx, customer)
	// if errInsert != nil {
	// 	return nil, errInsert
	// }

	return nil, nil
}
