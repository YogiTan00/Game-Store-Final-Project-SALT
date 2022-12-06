package item

import (
	"context"
	_repository "game-store-final-project/project/domain/repository"
)

type ItemHandler struct {
	ctx      context.Context
	repoItem _repository.ItemRepository
}

func NewItemHandler(ctx context.Context, repoItem _repository.ItemRepository) *ItemHandler {
	return &ItemHandler{
		ctx:      ctx,
		repoItem: repoItem,
	}
}
