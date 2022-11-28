package repository

import (
	"context"
	"game-store-final-project/project/domain/entity"
)

type TransactionRepository interface {
	GetAllTransaction(ctx context.Context) ([]*entity.Transaction, error)
}
