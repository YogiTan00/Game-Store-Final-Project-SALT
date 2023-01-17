package usecase

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
)

type TransactionUseCase interface {
	GetTransactionByID(ctx context.Context, id string) (*transaction.Transaction, error)
	GetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error)
	StoreTransaction(customer_id int, tanggal_pembelian string, voucher []string, items []*transaction.DTOItemPembelian) (*transaction.Transaction, error)
	GetAllTransactionByCustomerID(ctx context.Context, id string) ([]*transaction.Transaction, error)
}
