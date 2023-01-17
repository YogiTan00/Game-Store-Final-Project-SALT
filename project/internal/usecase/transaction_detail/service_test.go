package transaction_detail_test

import (
	"game-store-final-project/project/internal/usecase/transaction_detail"
	"game-store-final-project/project/test_data"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestTransactionDetailUseCaseInteractor_GetAllTransactionDetail(t *testing.T) {
	repoTransDMocks.On("GetAllTransactionDetail", mock.Anything).Return(test_data.GetTestDataCountTransactionDetail(3), (error)(nil))
	listTransD := transaction_detail.NewTransactionDetailUseCaseInteractor(repoTransDMocks, repoItemMocks)

	assert.NotNil(t, listTransD)
}
