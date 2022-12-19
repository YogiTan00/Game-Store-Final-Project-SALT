package transaction_handler_test

// func TestTransactionDetailHandler_GetAllTransactionDetailHandler(t *testing.T) {
// 	var (
// 		useCaseItem              = new(item.RepoItem)
// 		useCaseTransactionDetail = new(transaction_detail.RepoTransactionDetail)
// 	)

// 	useCaseTransactionDetail.On("GetAllTransactionDetail", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataCountTransactionDetail(5), (error)(nil))

// 	transactionDetailHandler := transaction_handler.NewTransactionDetailHandler(useCaseTransactionDetail, useCaseItem)

// 	req, err := http.NewRequest("GET", "/get-transaction-detail", nil)
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(transactionDetailHandler.GetAllTransactionDetailHandler)
// 	handler.ServeHTTP(rr, req)
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	assert.Nil(t, err)
// }

// func TestTransactionDetailHandler_GeAllTransactionDetailByIDHandler(t *testing.T) {
// 	var (
// 		useCaseItem              = new(item.RepoItem)
// 		useCaseTransactionDetail = new(transaction_detail.RepoTransactionDetail)
// 	)

// 	useCaseTransactionDetail.On("GetAllTransactionDetailByID", mock.Anything, mock.AnythingOfType("string")).Return(test_data.GetTestDataCountTransactionDetail(5), (error)(nil))

// 	transactionDetailHandler := transaction_handler.NewTransactionDetailHandler(useCaseTransactionDetail, useCaseItem)

// 	req, err := http.NewRequest("GET", "/get-transaction-detail/1", nil)
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(transactionDetailHandler.GeAllTransactionDetailByIDHandler)
// 	handler.ServeHTTP(rr, req)
// 	assert.Equal(t, http.StatusOK, rr.Code)
// 	assert.Nil(t, err)
// }
