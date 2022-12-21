package usecase

import (
	"context"
	"game-store-final-project/project/domain/entity/customer"
	"game-store-final-project/project/domain/entity/voucher"
)

type VoucherCase interface {
<<<<<<< HEAD
	UcStoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error
	UCGetVoucherByCustomerId(ctx context.Context, id string) (*customer.Customer, error)
=======
	StoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error
>>>>>>> 8b6c68ad914b0db0853cd7343d0c97cc30aee9d9
}
