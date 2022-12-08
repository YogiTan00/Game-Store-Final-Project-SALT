package usecase

import (
	"context"
	"game-store-final-project/project/domain/entity/item"
)

type ItemUseCase interface {
	UcGetAllItem(ctx context.Context) ([]*item.Item, error)
	UcGetItemByID(ctx context.Context, id string) (*item.Item, error)
	UcGetAllItemByID(ctx context.Context, id string) ([]*item.Item, error)
}
