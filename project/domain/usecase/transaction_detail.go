package usecase

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction_detail"
)

type TransactionDetailUseCase interface {
	GetAllTransactionDetail(ctx context.Context) ([]*transaction_detail.TransactionDetail, error)
	GetTransactionDetailByID(ctx context.Context, id string) (*transaction_detail.TransactionDetail, error)
	GetAllTransactionDetailByID(ctx context.Context, id string) ([]*transaction_detail.TransactionDetail, error)
}
