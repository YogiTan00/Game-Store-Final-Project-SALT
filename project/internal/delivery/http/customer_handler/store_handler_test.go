package customer_handler_test

import (
	"game-store-final-project/project/internal/delivery/http/customer_handler"
	"game-store-final-project/project/internal/usecase/customer"
	"game-store-final-project/project/test_data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	useCaseCustomer = new(customer.RepoCustomer)
)

// ini unit tes untuk usecase
func TestCustomerHandlerInteractor_StoreController(t *testing.T) {
	expected := &customer_handler.CustomerHandlerInteractor{CustomerUseCase: useCaseCustomer}

	customerHandler := customer_handler.NewUseCaseCustomerHandler(useCaseCustomer)

	assert.NotNil(t, customerHandler)
	assert.Equalf(t, expected, customerHandler, "NewUseCaseCustomerHandler(%v)", useCaseCustomer)
}

func TestCustomerHandlerInteractor_IndexController(t *testing.T) {
	useCaseCustomer.On("IndexCustomerWithTransaction", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataCustomer(), (error)(nil))

	customerHandler := customer_handler.NewUseCaseCustomerHandler(useCaseCustomer)
	req, err := http.NewRequest("GET", "/customer/list-trx/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(customerHandler.IndexController)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}
