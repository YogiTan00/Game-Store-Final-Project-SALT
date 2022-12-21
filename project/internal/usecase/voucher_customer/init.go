package voucher_customer

import (
	"game-store-final-project/project/domain/repository"
	"github.com/stretchr/testify/mock"
)

type RepoVoucherCustomer struct {
	mock.Mock
}

type VoucherCustomerUseCaseInteractor struct {
	repoVoucherCustomer repository.VoucherRepository
	repoCustomer        repository.CustomerRepository
}

func NewVoucherCustomerUseCaseInteractor(repoVoucherCustomer repository.VoucherRepository, repoCustomer repository.CustomerRepository) *VoucherCustomerUseCaseInteractor {
	return &VoucherCustomerUseCaseInteractor{
		repoVoucherCustomer: repoVoucherCustomer,
		repoCustomer:        repoCustomer,
	}
}
