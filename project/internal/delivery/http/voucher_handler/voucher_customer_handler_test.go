package voucher_handler_test

import (
	"game-store-final-project/project/internal/delivery/http/voucher_handler"
	customer2 "game-store-final-project/project/internal/usecase/customer"
	"game-store-final-project/project/internal/usecase/voucher_customer"
	"game-store-final-project/project/test_data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVoucherHandler_GetVoucherCustomerByIdHandler(t *testing.T) {
	var (
		useCaseVoucher  = new(voucher_customer.RepoVoucherCustomer)
		useCaseCustomer = new(customer2.RepoCustomer)
	)

	useCaseVoucher.On("GetVoucherByCustomerId", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataVoucherCustomer(3), (nil))

	voucherHandler := voucher_handler.NewVoucherHandler(useCaseVoucher, useCaseCustomer)

	req, err := http.NewRequest("GET", "/get-voucher/1", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(voucherHandler.GetVoucherCustomerByIdHandler)
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Nil(t, err)
}


