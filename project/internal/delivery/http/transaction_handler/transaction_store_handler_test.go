package transaction_handler_test

import (
	"game-store-final-project/project/internal/delivery/http/transaction_handler"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransactionHandler_StoreController(t *testing.T) {
	expected := &transaction_handler.TransactionHandlerInteractor{TransactionUseCase: useCaseTransaction}

	transactionHandler := transaction_handler.NewUsecaseTransactionHandler(useCaseTransaction)
	assert.NotNil(t, transactionHandler)
	assert.Equalf(t, expected, transactionHandler, "NewUsecaseTransactionHandler(%v)", useCaseTransaction)
}
