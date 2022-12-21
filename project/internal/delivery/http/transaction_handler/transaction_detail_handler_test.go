package transaction_handler_test

import (
	"game-store-final-project/project/internal/delivery/http/transaction_handler"
	"game-store-final-project/project/internal/usecase/item"
	"game-store-final-project/project/internal/usecase/transaction_detail"
	"game-store-final-project/project/test_data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTransactionDetailHandler_GetAllTransactionDetailHandler(t *testing.T) {
	var (
		useCaseItems             = new(item.RepoItem)
		useCaseTransactionDetail = new(transaction_detail.RepoTransactionDetail)
	)
	useCaseTransactionDetail.On("GetAllTransactionDetail", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataCountTransactionDetail(5), (error)(nil))
	transactionDetailHandler := transaction_handler.NewUseCaseTransactionDetailHandler(useCaseTransactionDetail, useCaseItems)

	req, err := http.NewRequest("GET", "/get-transaction-detail", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionDetailHandler.GetAllTransactionDetailHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}

func TestTransactionDetailHandler_GetTransactionDetailByIDHandler(t *testing.T) {
	var (
		useCaseItems             = new(item.RepoItem)
		useCaseTransactionDetail = new(transaction_detail.RepoTransactionDetail)
	)

	useCaseTransactionDetail.On("GetTransactionDetailByID", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataTransactionDetail(), (error)(nil))
	transactionDetailHandler := transaction_handler.NewUseCaseTransactionDetailHandler(useCaseTransactionDetail, useCaseItems)

	req, err := http.NewRequest("GET", "/get-transaction-detail/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionDetailHandler.GetTransactionDetailByIDHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}

func TestTransactionDetailHandler_GeAllTransactionDetailByIDHandler(t *testing.T) {
	var (
		useCaseItem              = new(item.RepoItem)
		useCaseTransactionDetail = new(transaction_detail.RepoTransactionDetail)
	)

	useCaseTransactionDetail.On("GetAllTransactionDetailByID", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataCountTransactionDetail(5), (error)(nil))
	transactionDetailHandler := transaction_handler.NewUseCaseTransactionDetailHandler(useCaseTransactionDetail, useCaseItem)

	req, err := http.NewRequest("GET", "/get-transaction-detail/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionDetailHandler.GetAllTransactionDetailByIDHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}
