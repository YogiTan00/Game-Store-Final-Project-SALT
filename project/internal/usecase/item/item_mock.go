package item

import (
	"context"
	"game-store-final-project/project/domain/entity/item"
)

func (r *RepoItem) GetAllItem(ctx context.Context) ([]*item.Item, error) {
	args := r.Called(ctx)
	return args.Get(0).([]*item.Item), args.Error(1)
}

func (r *RepoItem) GetItemByID(ctx context.Context, id string) (*item.Item, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(*item.Item), args.Error(1)
}
