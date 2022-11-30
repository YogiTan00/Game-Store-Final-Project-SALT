package transaction_handler

import (
	"context"
	_repository "game-store-final-project/project/domain/repository"
)

type TransactionHandler struct {
	ctx             context.Context
	repoTransaction _repository.TransactionRepository
}

type TransactionDetailHandler struct {
	ctx                   context.Context
	repoTransactionDetail _repository.TransactionDetailRepository
}

func NewTransactionHandler(ctx context.Context, repoTransaction _repository.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{
		ctx:             ctx,
		repoTransaction: repoTransaction,
	}
}

func NewTransactionDetailHandler(ctx context.Context, repoTransactionDetail _repository.TransactionDetailRepository) *TransactionDetailHandler {
	return &TransactionDetailHandler{
		ctx:                   ctx,
		repoTransactionDetail: repoTransactionDetail,
	}
}
