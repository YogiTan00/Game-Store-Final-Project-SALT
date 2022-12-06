package repository

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
)

type TransactionRepository interface {
	GetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error)
	StoreTransaction(ctx context.Context, dataTransaction *transaction.Transaction) error
}
