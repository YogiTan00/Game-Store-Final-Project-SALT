package transaction_handler

import (
	"context"
	_repository "game-store-final-project/project/domain/repository"
	"game-store-final-project/project/domain/usecase"
	transaction2 "game-store-final-project/project/internal/usecase/transaction"
	"game-store-final-project/project/internal/usecase/transaction_detail"
)

type TransactionHandler struct {
	useCaseTransaction    usecase.TransactionUseCase
	useCaseTransDetail    usecase.TransactionDetailUseCase
	repoTransaction       _repository.TransactionRepository
	repoTransactionDetail _repository.TransactionDetailRepository
	repoItem              _repository.ItemRepository
	repoVoucher           _repository.VoucherRepository
}

type TransactionDetailHandler struct {
	useCaseTransactionDetail usecase.TransactionDetailUseCase
	repoTransactionDetail    _repository.TransactionDetailRepository
	repoItem                 _repository.ItemRepository
}

type TransactionHandlerInteractor struct {
	TransactionUseCase usecase.TransactionUseCase
}

func NewTransactionHandler(repoTransaction _repository.TransactionRepository, repoTransactionDetail _repository.TransactionDetailRepository, repoItem _repository.ItemRepository, repoVoucher _repository.VoucherRepository) *TransactionHandler {
	var (
		ctx = context.Background()
	)
	useCaseTransaction := transaction2.NewTransactionUseCaseInteractor(ctx, repoTransaction, repoItem, repoVoucher, repoTransactionDetail)
	useCaseTransDetail := transaction_detail.NewTransactionDetailUseCaseInteractor(repoTransactionDetail)
	return &TransactionHandler{
		useCaseTransaction:    useCaseTransaction,
		useCaseTransDetail:    useCaseTransDetail,
		repoTransaction:       repoTransaction,
		repoTransactionDetail: repoTransactionDetail,
		repoItem:              repoItem,
		repoVoucher:           repoVoucher,
	}
}

func NewTransactionDetailHandler(repoTransactionDetail _repository.TransactionDetailRepository, repoItem _repository.ItemRepository) *TransactionDetailHandler {
	transactionDetail := transaction_detail.NewTransactionDetailUseCaseInteractor(repoTransactionDetail)
	return &TransactionDetailHandler{
		useCaseTransactionDetail: transactionDetail,
		repoTransactionDetail:    repoTransactionDetail,
		repoItem:                 repoItem,
	}
}

func NewUsecaseTransactionHandler(usecaseImplement usecase.TransactionUseCase) *TransactionHandlerInteractor {
	return &TransactionHandlerInteractor{TransactionUseCase: usecaseImplement}
}
