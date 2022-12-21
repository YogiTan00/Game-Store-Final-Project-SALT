package voucher

import (
	"context"
	"game-store-final-project/project/domain/entity/voucher"
)

func (v VoucherUseCaseInteractor) StoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error {
	Item := v.repoVoucher.StoreVoucher(ctx, dataVoucher)

	return Item
}
