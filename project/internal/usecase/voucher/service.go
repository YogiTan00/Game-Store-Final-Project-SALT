package voucher

import (
	"context"
	"game-store-final-project/project/domain/entity/voucher"
)

func (v VoucherUseCaseInteractor) UcStoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error {
	Item := v.repoVoucher.StoreVoucher(ctx, dataVoucher)

	return Item
}
