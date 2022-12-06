package repository

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
)

type TransactionDetailRepository interface {
	GetAllTransactionDetail(ctx context.Context) ([]*transaction.TransactionDetail, error)
	GetTransactionDetailByID(ctx context.Context, id string) (*transaction.TransactionDetail, error)
	GetAllTransactionDetailByID(ctx context.Context, id string) ([]*transaction.TransactionDetail, error)
}
