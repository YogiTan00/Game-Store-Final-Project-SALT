package voucher

import (
	"context"
	"game-store-final-project/project/domain/entity/voucher"
)

func (r *RepoVoucher) StoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error {
	args := r.Called(ctx, dataVoucher)
	return args.Error(0)
}
