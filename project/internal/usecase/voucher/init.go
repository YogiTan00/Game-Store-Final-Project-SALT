package voucher

import (
	"game-store-final-project/project/domain/repository"
	"github.com/stretchr/testify/mock"
)

type RepoVoucher struct {
	mock.Mock
}

type VoucherUseCaseInteractor struct {
	repoVoucher repository.VoucherRepository
}

func NewVoucherUseCaseInteractor(repoVoucher repository.VoucherRepository) *VoucherUseCaseInteractor {
	return &VoucherUseCaseInteractor{
		repoVoucher: repoVoucher,
	}
}
