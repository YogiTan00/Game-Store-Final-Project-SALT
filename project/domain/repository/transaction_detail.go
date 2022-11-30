package repository

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
)

type TransactionDetailRepository interface {
	GetAllTransactionDetail(ctx context.Context) ([]*transaction.TransactionDetail, error)
}
