package transaction

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
)

func (r *RepoTransaction) GetTransactionByID(ctx context.Context, id string) (*transaction.Transaction, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(*transaction.Transaction), args.Error(1)
}

func (r *RepoTransaction) GetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error) {
	args := r.Called(ctx)
	return args.Get(0).([]*transaction.Transaction), args.Error(1)
}

func (r *RepoTransaction) StoreTransaction(customer_id int, tanggal_pembelian string, voucher []string, items []*transaction.DTOItemPembelian) (*transaction.Transaction, error) {
	args := r.Called(customer_id, tanggal_pembelian, voucher, items)
	return args.Get(0).(*transaction.Transaction), args.Error(0)
}

func (r *RepoTransaction) GetAllTransactionByCustomerID(ctx context.Context, id string) ([]*transaction.Transaction, error) {
	args := r.Called(ctx, id)
	return args.Get(0).([]*transaction.Transaction), args.Error(1)
}
