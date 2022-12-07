package item

import (
	_repository "game-store-final-project/project/domain/repository"
	"game-store-final-project/project/domain/usecase"
	"game-store-final-project/project/internal/usecase/item"
)

type ItemHandler struct {
	itemUseCase usecase.ItemUseCase
	repoItem    _repository.ItemRepository
}

func NewItemHandler(repoItem _repository.ItemRepository) *ItemHandler {
	itemUseCase := item.NewItemUseCaseInteractor(repoItem)
	return &ItemHandler{
		itemUseCase: itemUseCase,
		repoItem:    repoItem,
	}
}
