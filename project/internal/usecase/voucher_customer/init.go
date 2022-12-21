package voucher_customer

import (
	"game-store-final-project/project/domain/repository"
	"github.com/stretchr/testify/mock"
)

type RepoVoucherCustomerById struct {
	mock.Mock
}

type VoucherCustomerUseCaseInteractor struct {
	repoVoucherCustomer repository.VoucherRepository
}

func NewVoucherCustomerUseCaseInteractor(repoVoucherCustomerById repository.VoucherRepository) *VoucherCustomerUseCaseInteractor {
	return &VoucherCustomerUseCaseInteractor{
		repoVoucherCustomer: repoVoucherCustomerById,
	}
}
