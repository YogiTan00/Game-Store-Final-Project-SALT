package voucher

import (
	"context"
	"game-store-final-project/project/domain/entity/voucher"
)

func (r *RepoVoucher) StoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error {
	args := r.Called(ctx, dataVoucher)
	return args.Error(0)
}

//func (r *RepoVoucher) GetVoucherByCustomerId(ctx context.Context, id string) (*customer.Customer, error) {
//	args := r.Called(ctx, id)
//	return args.Get(0).(*customer.Customer), args.Error(1)
//}
