package transaction

import (
	"context"
	"game-store-final-project/project/domain/repository"
)

type TransactionUseCaseInteractor struct {
	ctx             context.Context
	repoTransaction repository.TransactionRepository
	repoItem        repository.ItemRepository
}

// ini ngeimplement domain usecase
func NewTransactionUseCaseInteractor(ctx context.Context, repoImplement repository.TransactionRepository, repoImplementItem repository.ItemRepository) *TransactionUseCaseInteractor {
	return &TransactionUseCaseInteractor{
		ctx:             ctx,
		repoTransaction: repoImplement,
		repoItem:        repoImplementItem,
	}
}
