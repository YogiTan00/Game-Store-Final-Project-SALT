package voucher

import (
	"context"
	"game-store-final-project/project/domain/entity/voucher"
)

func (v *VoucherUseCaseInteractor) StoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error {
	Item := v.repoVoucher.StoreVoucher(ctx, dataVoucher)

	return Item
}

//func (v *VoucherUseCaseInteractor) GetVoucherByCustomerId(ctx context.Context, id string) (*customer.Customer, error) {
//	listVoucherCustomer, err := v.repoVoucher.GetVoucherByCustomerId(ctx, id)
//	if err != nil {
//		return nil, err
//	}
//
//	return listVoucherCustomer, nil
//}
