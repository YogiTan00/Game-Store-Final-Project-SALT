package voucher_handler

import (
	"game-store-final-project/project/domain/repository"
	"game-store-final-project/project/domain/usecase"
	"game-store-final-project/project/internal/usecase/voucher"
)

type VoucherHandler struct {
	useCaseVoucher usecase.VoucherCase
	repoVoucher    repository.VoucherRepository
}

func NewVoucherHandler(repoVoucher repository.VoucherRepository) *VoucherHandler {
	voucherUseCase := voucher.NewVoucherUseCaseInteractor(repoVoucher)
	return &VoucherHandler{
		useCaseVoucher: voucherUseCase,
		repoVoucher:    repoVoucher,
	}
}
