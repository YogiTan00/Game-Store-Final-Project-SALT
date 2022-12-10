package transaction

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
)

func (r *RepoTransaction) GetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error) {
	args := r.Called(ctx)
	return args.Get(0).([]*transaction.Transaction), args.Error(1)
}

func (r *RepoTransaction) StoreTransaction(ctx context.Context, dataTransaction *transaction.Transaction) error {
	args := r.Called(ctx, dataTransaction)
	return args.Error(0)
}

func (r *RepoTransaction) GetAllTransactionByID(ctx context.Context, id string) ([]*transaction.Transaction, error) {
	args := r.Called(ctx, id)
	return args.Get(0).([]*transaction.Transaction), args.Error(1)
}
