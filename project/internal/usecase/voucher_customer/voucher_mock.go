package voucher_customer

import (
	"context"
	"game-store-final-project/project/domain/entity/customer"
)

func (r *RepoVoucherCustomer) GetVoucherByCustomerId(ctx context.Context, id string) (*customer.Customer, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(*customer.Customer), args.Error(1)
}
