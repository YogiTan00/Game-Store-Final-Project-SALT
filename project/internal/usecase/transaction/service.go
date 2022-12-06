package transaction

import (
	"fmt"
	"game-store-final-project/project/domain/entity/transaction"
)

func (cu *TransactionUseCaseInteractor) StoreTransaction(customer_id int, tanggal_pembelian string, voucher string, items []*transaction.DTOItemPembelian) (*transaction.Transaction, error) {
	// build data to entity
	transaction, err := transaction.NewTransaction(transaction.DTOTransaction{
		VoucherCustomerId: 1,
		CustomerId:        customer_id,
		CodeTransaction:   "",
		Tanggalpembelian:  tanggal_pembelian,
		Total:             0,
		HargaDiscount:     0,
		TotalHarga:        0,
		ItemPembelian:     items,
	})

	if err != nil { // error build
		return nil, err
	}

	fmt.Println(transaction)

	// errInsert := cu.RepoCustomer.StoreCustomer(cu.ctx, customer)
	// if errInsert != nil {
	// 	return nil, errInsert
	// }

	return nil, nil
}
