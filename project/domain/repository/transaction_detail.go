package repository

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction_detail"
)

type TransactionDetailRepository interface {
	GetAllTransactionDetail(ctx context.Context) ([]*transaction_detail.TransactionDetail, error)
	GetTransactionDetailByID(ctx context.Context, id string) (*transaction_detail.TransactionDetail, error)
	GetAllTransactionDetailByID(ctx context.Context, id string) ([]*transaction_detail.TransactionDetail, error)
	StoreTransactionDetail(ctx context.Context, trx_id int64, detail *transaction_detail.TransactionDetail) error
}
