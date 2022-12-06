package transaction

import (
	"game-store-final-project/project/domain/entity/transaction"
)

func (cu *TransactionUseCaseInteractor) StoreTransaction(customer_id int, tanggal_pembelian string, voucher string, items []*transaction.DTOItemPembelian) (*transaction.Transaction, error) {
	/*
		Rule:1
		cek dapet diskon atau tidak,
		jika iya generate ke table discount dan return datanya

		Rule:2
		cek produknya ready atau tidak, jika iya return datanya
	*/

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

	errInsert := cu.repoTransaction.StoreTransaction(cu.ctx, transaction)
	if errInsert != nil {
		return nil, errInsert
	}

	return transaction, nil
}
