package usecase

import (
	"context"
	"game-store-final-project/project/domain/entity/item"
)

type ItemUseCase interface {
	GetAllItem(ctx context.Context) ([]*item.Item, error)
	GetItemByID(ctx context.Context, id string) (*item.Item, error)
}
