package transaction_handler_test

import (
	"game-store-final-project/project/internal/delivery/http/transaction_handler"
	"game-store-final-project/project/internal/usecase/item"
	"game-store-final-project/project/internal/usecase/transaction"
	"game-store-final-project/project/internal/usecase/transaction_detail"
	"game-store-final-project/project/internal/usecase/voucher"
	"game-store-final-project/project/test_data"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	useCaseItem        = new(item.RepoItem)
	useCaseTransaction = new(transaction.RepoTransaction)
	useCaseTransDetail = new(transaction_detail.RepoTransactionDetail)
	useCaseVoucher     = new(voucher.RepoVoucher)
)

func TestTransactionHandler_GetAllTransactionHandler(t *testing.T) {

	useCaseTransaction.On("GetAllTransaction", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataCountTransaction(5), (error)(nil))

	transactionDetailHandler := transaction_handler.NewTransactionHandler(useCaseTransaction, useCaseTransDetail, useCaseItem, useCaseVoucher)

	req, err := http.NewRequest("GET", "/get-transaction", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionDetailHandler.GetAllTransactionHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}

func TestTransactionHandler_GetAllTransactionByIDHandler(t *testing.T) {

	useCaseTransaction.On("GetAllTransactionByCustomerID", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataCountTransaction(5), (error)(nil))

	transactionDetailHandler := transaction_handler.NewTransactionHandler(useCaseTransaction, useCaseTransDetail, useCaseItem, useCaseVoucher)

	req, err := http.NewRequest("GET", "/get-transaction/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(transactionDetailHandler.GetAllTransactionByIDHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}
