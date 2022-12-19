package usecase

import (
	"context"
	"game-store-final-project/project/domain/entity/voucher"
)

type VoucherCase interface {
	StoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error
}
