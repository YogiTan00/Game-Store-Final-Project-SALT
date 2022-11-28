package transaction_handler

import (
	"context"
	_repository "game-store-final-project/project/domain/repository"
)

type TransactionHandler struct {
	ctx             context.Context
	repoTransaction _repository.TransactionRepository
}

func NewCustomerHandler(ctx context.Context, repotransaction _repository.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{
		ctx:             ctx,
		repoTransaction: repotransaction,
	}
}
