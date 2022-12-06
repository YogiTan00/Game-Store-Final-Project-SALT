package transaction

import (
	"context"
	"game-store-final-project/project/domain/repository"
)

type TransactionUseCaseInteractor struct {
	ctx             context.Context
	repoTransaction repository.TransactionRepository
}

// ini ngeimplement domain usecase
func NewTransactionUseCaseInteractor(ctx context.Context, repoImplement repository.TransactionRepository) *TransactionUseCaseInteractor {
	return &TransactionUseCaseInteractor{ctx: ctx, repoTransaction: repoImplement}
}
