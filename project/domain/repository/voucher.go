package repository

import (
	"context"
	"game-store-final-project/project/domain/entity/voucher"
)

type VoucherRepository interface {
	StoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error
}
