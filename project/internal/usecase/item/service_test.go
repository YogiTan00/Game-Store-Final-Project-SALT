package item_test

import (
	"game-store-final-project/project/internal/usecase/item"
	"game-store-final-project/project/test_data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestItemUseCaseInteractor_GetAllItem(t *testing.T) {
	repoItemMocks.On("GetAllitem", mock.Anything).Return(test_data.GetTestDataCountCustomer(5), (error)(nil))
	listItem := item.NewItemUseCaseInteractor(repoItemMocks)

	assert.NotNil(t, listItem)
}
