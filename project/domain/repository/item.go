package repository

import (
	"context"
	"game-store-final-project/project/domain/entity/item"
)

type ItemRepository interface {
	GetAllItem(ctx context.Context) ([]*item.Item, error)
	GetItemByID(ctx context.Context, id string) (*item.Item, error)
}
