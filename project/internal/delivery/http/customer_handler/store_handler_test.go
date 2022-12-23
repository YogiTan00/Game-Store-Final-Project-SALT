package customer_handler_test

import (
	"game-store-final-project/project/internal/delivery/http/customer_handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ini unit tes untuk usecase
func TestCustomerHandlerInteractor_StoreController(t *testing.T) {
	expected := &customer_handler.CustomerHandlerInteractor{CustomerUseCase: useCaseCustomer}
	customerHandler := customer_handler.NewUseCaseCustomerHandler(useCaseCustomer)
	assert.NotNil(t, customerHandler)
	assert.Equalf(t, expected, customerHandler, "NewUseCaseCustomerHandler(%v)", useCaseCustomer)
}
