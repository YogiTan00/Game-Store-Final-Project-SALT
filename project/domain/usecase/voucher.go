package usecase

import (
	"context"
	"game-store-final-project/project/domain/entity/customer"
	"game-store-final-project/project/domain/entity/voucher"
)

type VoucherCase interface {
	UcStoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error
	UCGetVoucherByCustomerId(ctx context.Context, id string) (*customer.Customer, error)
}
