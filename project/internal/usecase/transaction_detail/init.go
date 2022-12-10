package transaction_detail

import (
	"game-store-final-project/project/domain/repository"
	"github.com/stretchr/testify/mock"
)

type RepoTransactionDetail struct {
	mock.Mock
}

type TransactionDetailUseCaseInteractor struct {
	repoTransactionDetail repository.TransactionDetailRepository
}

func NewTransactionDetailUseCaseInteractor(repoTransactionDetail repository.TransactionDetailRepository) *TransactionDetailUseCaseInteractor {
	return &TransactionDetailUseCaseInteractor{
		repoTransactionDetail: repoTransactionDetail,
	}
}
