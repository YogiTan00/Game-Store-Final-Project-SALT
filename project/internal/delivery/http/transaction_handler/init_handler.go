package transaction_handler

import (
	"game-store-final-project/project/domain/usecase"
)

type TransactionHandler struct {
	useCaseTransaction usecase.TransactionUseCase
	useCaseTransDetail usecase.TransactionDetailUseCase
	useCaseItem        usecase.ItemUseCase
	useCaseVoucher     usecase.VoucherUseCase
}

type TransactionDetailHandler struct {
	useCaseTransactionDetail usecase.TransactionDetailUseCase
	useCaseItem              usecase.ItemUseCase
}

type TransactionHandlerInteractor struct {
	TransactionUseCase usecase.TransactionUseCase
}

func NewUseCaseTransactionHandler(
	useCaseTransaction usecase.TransactionUseCase,
	useCaseTransDetail usecase.TransactionDetailUseCase,
	useCaseItem usecase.ItemUseCase,
	useCaseVoucher usecase.VoucherUseCase,
) *TransactionHandler {

	return &TransactionHandler{
		useCaseTransaction: useCaseTransaction,
		useCaseTransDetail: useCaseTransDetail,
		useCaseItem:        useCaseItem,
		useCaseVoucher:     useCaseVoucher,
	}
}

func NewUseCaseTransactionDetailHandler(useCaseTransactionDetail usecase.TransactionDetailUseCase, useCaseItem usecase.ItemUseCase) *TransactionDetailHandler {
	return &TransactionDetailHandler{
		useCaseTransactionDetail: useCaseTransactionDetail,
		useCaseItem:              useCaseItem,
	}
}

func NewUsecaseTransactionHandler(usecaseImplement usecase.TransactionUseCase) *TransactionHandlerInteractor {
	return &TransactionHandlerInteractor{TransactionUseCase: usecaseImplement}
}
