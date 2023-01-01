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
	repoItem              repository.ItemRepository
}

func NewTransactionDetailUseCaseInteractor(repoTransactionDetail repository.TransactionDetailRepository, repoItem repository.ItemRepository) *TransactionDetailUseCaseInteractor {
	return &TransactionDetailUseCaseInteractor{
		repoTransactionDetail: repoTransactionDetail,
		repoItem:              repoItem,
	}
}
