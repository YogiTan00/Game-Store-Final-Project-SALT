package item_handler_test

import (
	item2 "game-store-final-project/project/internal/delivery/http/item_handler"
	"game-store-final-project/project/internal/usecase/item"
	"game-store-final-project/project/test_data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestItemHandler_GetAllItem(t *testing.T) {
	var (
		useCaseItem = new(item.RepoItem)
	)

	useCaseItem.On("GetAllItem", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataCountItem(5), (error)(nil))

	itemHandler := item2.NewuseCaseItemHandler(useCaseItem)

	req, err := http.NewRequest("GET", "/get-item", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(itemHandler.GetAllItemHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}

func TestItemHandler_GetItemByID(t *testing.T) {
	var (
		useCaseItem = new(item.RepoItem)
	)

	useCaseItem.On("GetItemByID", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataItem(), (nil))

	itemHandler := item2.NewuseCaseItemHandler(useCaseItem)

	req, err := http.NewRequest("GET", "/get-item/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(itemHandler.GetItemByIDHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}
