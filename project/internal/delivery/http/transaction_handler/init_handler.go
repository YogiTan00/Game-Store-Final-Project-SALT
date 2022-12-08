package transaction_handler

import (
	"context"
	_repository "game-store-final-project/project/domain/repository"
	"game-store-final-project/project/domain/usecase"
	"game-store-final-project/project/internal/usecase/transaction_detail"
)

type TransactionHandler struct {
	ctx             context.Context
	repoTransaction _repository.TransactionRepository
}

type TransactionDetailHandler struct {
	useCaseTransactionDetail usecase.TransactionDetailUseCase
	repoTransactionDetail    _repository.TransactionDetailRepository
	repoItem                 _repository.ItemRepository
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

func NewTransactionDetailHandler(repoTransactionDetail _repository.TransactionDetailRepository, repoItem _repository.ItemRepository) *TransactionDetailHandler {
	var (
		ctx = context.Background()
	)
	transactionDetail := transaction_detail.NewTransactionDetailUseCaseInteractor(ctx, repoTransactionDetail)
	return &TransactionDetailHandler{
		useCaseTransactionDetail: transactionDetail,
		repoTransactionDetail:    repoTransactionDetail,
		repoItem:                 repoItem,
	}
}

func NewUsecaseTransactionHandler(usecaseImplement usecase.TransactionUseCase) *TransactionHandlerInteractor {
	return &TransactionHandlerInteractor{TransactionUseCase: usecaseImplement}
}
