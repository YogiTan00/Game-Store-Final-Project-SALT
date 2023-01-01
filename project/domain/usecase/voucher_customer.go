package usecase

import (
	"context"
	"game-store-final-project/project/domain/entity/customer"
)

type VoucherCustomerUseCase interface {
	GetVoucherByCustomerId(ctx context.Context, id string) (*customer.Customer, error)
}
