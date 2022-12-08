package usecase

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction_detail"
)

type TransactionDetailUseCase interface {
	UcGetAllTransactionDetail(ctx context.Context) ([]*transaction_detail.TransactionDetail, error)
	UcGetTransactionDetailByID(ctx context.Context, id string) (*transaction_detail.TransactionDetail, error)
	UcGetAllTransactionDetailByID(ctx context.Context, id string) ([]*transaction_detail.TransactionDetail, error)
}
