package usecase

import (
	"game-store-final-project/project/domain/entity/transaction"
)

type TransactionUseCase interface {
	StoreTransaction(customer_id int, tanggal_pembelian string, voucher string, items []*transaction.DTODetailTransaction) (*transaction.Transaction, error)
}
