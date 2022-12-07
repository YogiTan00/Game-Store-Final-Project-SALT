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
