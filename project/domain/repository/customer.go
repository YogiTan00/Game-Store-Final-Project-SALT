package repository

import (
	"context"
	"game-store-final-project/project/domain/entity"
)

type CustomerRepository interface {
	StoreCustomer(ctx context.Context, dataArticle *entity.Customer) error
}
