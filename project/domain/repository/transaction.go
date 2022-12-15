package repository

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
)

type TransactionRepository interface {
	GetTransactionByID(ctx context.Context, id string) (*transaction.Transaction, error)
	GetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error)
	StoreTransaction(ctx context.Context, dataTransaction *transaction.Transaction) (int64, error)
	GetAllTransactionByCustomerID(ctx context.Context, id string) ([]*transaction.Transaction, error)
}
