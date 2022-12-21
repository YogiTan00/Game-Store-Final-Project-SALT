package transaction

import (
	"context"
	"game-store-final-project/project/domain/repository"

	"github.com/stretchr/testify/mock"
)

type RepoTransaction struct {
	mock.Mock
}

type TransactionUseCaseInteractor struct {
	ctx                   context.Context
	repoTransaction       repository.TransactionRepository
	repoItem              repository.ItemRepository
	repoVoucher           repository.VoucherRepository
	repoTransactionDetail repository.TransactionDetailRepository
	repoCustomer          repository.CustomerRepository
}

// ini ngeimplement domain usecase
func NewTransactionUseCaseInteractor(
	ctx context.Context,
	repoImplement repository.TransactionRepository,
	repoImplementItem repository.ItemRepository,
	repoImplementVoucher repository.VoucherRepository,
	repoImplementTrxDetail repository.TransactionDetailRepository,
	repoImplementCustomer repository.CustomerRepository,
) *TransactionUseCaseInteractor {
	return &TransactionUseCaseInteractor{
		ctx:                   ctx,
		repoTransaction:       repoImplement,
		repoItem:              repoImplementItem,
		repoVoucher:           repoImplementVoucher,
		repoTransactionDetail: repoImplementTrxDetail,
		repoCustomer:          repoImplementCustomer,
	}
}
