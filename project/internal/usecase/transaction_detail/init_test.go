package transaction_detail_test

import (
	item2 "game-store-final-project/project/internal/usecase/item"
	"game-store-final-project/project/internal/usecase/transaction_detail"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	repoItemMocks   = new(item2.RepoItem)
	repoTransDMocks = new(transaction_detail.RepoTransactionDetail)
)

func TestNewTransactionDetailUseCaseInteractor(t *testing.T) {
	transD := transaction_detail.NewTransactionDetailUseCaseInteractor(repoTransDMocks, repoItemMocks)
	assert.NotNil(t, transD)
}
