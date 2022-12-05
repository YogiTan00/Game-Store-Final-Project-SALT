package repository

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
)

type TransactionRepository interface {
	GetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error)
	GetAllTransactionByID(ctx context.Context, id string) ([]*transaction.Transaction, error)
}
