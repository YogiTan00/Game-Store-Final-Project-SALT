package voucher_customer

import (
	"context"
	"game-store-final-project/project/domain/entity/customer"
)

func (uc *VoucherCustomerUseCaseInteractor) GetVoucherByCustomerId(ctx context.Context, id string) (*customer.Customer, error) {
	listVoucherCustomer, err := uc.repoVoucherCustomer.GetVoucherByCustomerId(ctx, id)
	if err != nil {
		return nil, err
	}

	return listVoucherCustomer, nil
}
