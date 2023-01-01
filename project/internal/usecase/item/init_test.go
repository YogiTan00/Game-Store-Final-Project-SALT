package item_test

import (
	item2 "game-store-final-project/project/internal/usecase/item"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	repoItemMocks = new(item2.RepoItem)
)

func TestNewItemUseCaseInteractor(t *testing.T) {
	item := item2.NewItemUseCaseInteractor(repoItemMocks)
	assert.NotNil(t, item)
}
