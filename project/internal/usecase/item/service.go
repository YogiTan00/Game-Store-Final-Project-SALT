package item

import (
	"context"
	"game-store-final-project/project/domain/entity/item"
)

func (uc *ItemUseCaseInteractor) GetAllItem(ctx context.Context) ([]*item.Item, error) {
	listItem, err := uc.repoItem.GetAllItem(ctx)
	if err != nil {
		return nil, err
	}

	return listItem, nil
}

func (uc *ItemUseCaseInteractor) GetItemByID(ctx context.Context, id string) (*item.Item, error) {
	listItem, err := uc.repoItem.GetItemByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return listItem, nil
}
