package customer_handler_test

import (
	"game-store-final-project/project/internal/delivery/http/customer_handler"
	"game-store-final-project/project/internal/usecase/customer"
	"game-store-final-project/project/test_data"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	useCaseCustomer = new(customer.RepoCustomer)
)

func TestCustomerHandlerInteractor_IndexController(t *testing.T) {
	useCaseCustomer.On("IndexCustomerWithTransaction", mock.Anything, mock.AnythingOfType("string")).Return(test_data.DataCustomerWithTransaction(), (error)(nil))

	customerHandler := customer_handler.NewUseCaseCustomerHandler(useCaseCustomer)
	req, err := http.NewRequest("GET", "/customer/list-trx/3241203322112233", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(customerHandler.IndexController)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}
