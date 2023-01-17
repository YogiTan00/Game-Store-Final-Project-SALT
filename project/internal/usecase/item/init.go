package item

import (
	"game-store-final-project/project/domain/repository"
	"github.com/stretchr/testify/mock"
)

type RepoItem struct {
	mock.Mock
}

type ItemUseCaseInteractor struct {
	repoItem repository.ItemRepository
}

func NewItemUseCaseInteractor(repoItem repository.ItemRepository) *ItemUseCaseInteractor {
	return &ItemUseCaseInteractor{
		repoItem: repoItem,
	}
}
