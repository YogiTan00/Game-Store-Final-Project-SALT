package voucher_handler

import (
	"game-store-final-project/project/domain/usecase"
)

type VoucherHandler struct {
	useCaseVoucher  usecase.VoucherCustomerUseCase
	useCaseCustomer usecase.CustomerUseCase
}

func NewVoucherHandler(useCaseVoucher usecase.VoucherCustomerUseCase, useCaseCustomer usecase.CustomerUseCase) *VoucherHandler {
	return &VoucherHandler{
		useCaseVoucher:  useCaseVoucher,
		useCaseCustomer: useCaseCustomer,
	}
}
