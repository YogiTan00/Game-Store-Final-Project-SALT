package transaction

import (
	"context"
	"game-store-final-project/project/domain/repository"
	"github.com/stretchr/testify/mock"
)

type RepoTransaction struct {
	mock.Mock
}

type TransactionUseCaseInteractor struct {
	ctx             context.Context
	repoTransaction repository.TransactionRepository
}

// ini ngeimplement domain usecase
func NewTransactionUseCaseInteractor(ctx context.Context, repoTransaction repository.TransactionRepository) *TransactionUseCaseInteractor {
	return &TransactionUseCaseInteractor{
		ctx:             ctx,
		repoTransaction: repoTransaction,
	}
}
