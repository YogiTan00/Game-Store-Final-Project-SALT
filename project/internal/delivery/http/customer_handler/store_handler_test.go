package customer_handler_test

import (
	"net/http/httptest"
)

var (
	req = httptest.NewRequest("GET", "http://example.com/foo", nil)
	w   = httptest.NewRecorder()
)

// func TestHandlerStoreCustomer(t *testing.T) {
// 	useCasePhoneMocks.On("GetPhone").Return(testdata.GenerateListPhoneByEntity(), nil)

// 	handlerInteractorCustomer := customer_handler.NewCustomerHandler(useCasePhoneMocks)
// 	handlerInteractorCustomer.Store(w, req)
// 	resp := w.Result()
// 	body, _ := io.ReadAll(resp.Body)

// 	assert.Equal(t, http.StatusOK, resp.StatusCode)
// 	assert.NotNil(t, string(body))
// }
