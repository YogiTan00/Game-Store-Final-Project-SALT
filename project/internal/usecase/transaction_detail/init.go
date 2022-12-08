package transaction_detail

import (
	"context"
	"game-store-final-project/project/domain/repository"
	"github.com/stretchr/testify/mock"
)

type RepoTransactionDetail struct {
	mock.Mock
}

type TransactionDetailUseCaseInteractor struct {
	ctx                   context.Context
	repoTransactionDetail repository.TransactionDetailRepository
}

func NewTransactionDetailUseCaseInteractor(ctx context.Context, repoTransactionDetail repository.TransactionDetailRepository) *TransactionDetailUseCaseInteractor {
	return &TransactionDetailUseCaseInteractor{
		ctx:                   ctx,
		repoTransactionDetail: repoTransactionDetail,
	}
}
