package item_handler

import (
	"game-store-final-project/project/domain/usecase"
)

type ItemHandler struct {
	useCaseItem usecase.ItemUseCase
}

func NewuseCaseItemHandler(useCaseItem usecase.ItemUseCase) *ItemHandler {
	return &ItemHandler{
		useCaseItem: useCaseItem,
	}
}
