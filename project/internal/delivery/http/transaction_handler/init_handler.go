package transaction_handler

import (
	"context"
	_repository "game-store-final-project/project/domain/repository"
	"game-store-final-project/project/domain/usecase"
)

type TransactionHandler struct {
	ctx             context.Context
	repoTransaction _repository.TransactionRepository
}

type TransactionDetailHandler struct {
	ctx                   context.Context
	repoTransactionDetail _repository.TransactionDetailRepository
}

type TransactionHandlerInteractor struct {
	TransactionUseCase usecase.TransactionUseCase
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

func NewUsecaseTransactionHandler(usecaseImplement usecase.TransactionUseCase) *TransactionHandlerInteractor {
	return &TransactionHandlerInteractor{TransactionUseCase: usecaseImplement}
}
