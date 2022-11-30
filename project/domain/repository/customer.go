package repository

import (
	"context"
	"game-store-final-project/project/domain/entity/customer"
)

type CustomerRepository interface {
	StoreCustomer(ctx context.Context, dataCustomer *customer.Customer) error
}
