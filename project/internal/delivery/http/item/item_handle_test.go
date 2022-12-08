package item

import (
	"game-store-final-project/project/internal/usecase/item"
	"game-store-final-project/project/test_data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItemHandler_UcGetAllItem(t *testing.T) {
	var (
		useCaseItem = new(item.RepoItem)
	)

	useCaseItem.On("GetAllItem", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataCountItem(5), (error)(nil))

	itemHandler := NewItemHandler(useCaseItem)

	req, err := http.NewRequest("GET", "/get-item", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(itemHandler.GetAllItem)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}

func TestItemHandler_UcGetItemByID(t *testing.T) {
	var (
		useCaseItem = new(item.RepoItem)
	)

	useCaseItem.On("GetItemByID", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataItem(), (error)(nil))

	itemHandler := NewItemHandler(useCaseItem)

	req, err := http.NewRequest("GET", "/get-item/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(itemHandler.GetItemByID)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}
