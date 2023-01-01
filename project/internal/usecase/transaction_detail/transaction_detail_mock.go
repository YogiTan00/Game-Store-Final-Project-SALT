package transaction_detail

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction_detail"
)

func (r *RepoTransactionDetail) GetAllTransactionDetail(ctx context.Context) ([]*transaction_detail.TransactionDetail, error) {
	args := r.Called(ctx)
	return args.Get(0).([]*transaction_detail.TransactionDetail), args.Error(1)
}

func (r *RepoTransactionDetail) GetTransactionDetailByID(ctx context.Context, id string) (*transaction_detail.TransactionDetail, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(*transaction_detail.TransactionDetail), args.Error(1)
}

func (r *RepoTransactionDetail) GetAllTransactionDetailByID(ctx context.Context, id string) ([]*transaction_detail.TransactionDetail, error) {
	args := r.Called(ctx, id)
	return args.Get(0).([]*transaction_detail.TransactionDetail), args.Error(1)
}

func (r *RepoTransactionDetail) StoreTransactionDetail(ctx context.Context, trx_id int64, detail *transaction_detail.TransactionDetail) error {
	args := r.Called(ctx, trx_id, detail)
	return args.Error(0)
}
