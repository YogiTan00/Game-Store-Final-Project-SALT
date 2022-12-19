package transaction_detail

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction_detail"
)

func (uc *TransactionDetailUseCaseInteractor) GetAllTransactionDetail(ctx context.Context) ([]*transaction_detail.TransactionDetail, error) {
	listTransactionDetail, err := uc.repoTransactionDetail.GetAllTransactionDetail(ctx)
	if err != nil {
		return nil, err
	}

	return listTransactionDetail, nil
}

func (uc *TransactionDetailUseCaseInteractor) GetTransactionDetailByID(ctx context.Context, id string) (*transaction_detail.TransactionDetail, error) {
	listTransactionDetail, err := uc.repoTransactionDetail.GetTransactionDetailByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return listTransactionDetail, nil
}

func (uc *TransactionDetailUseCaseInteractor) GetAllTransactionDetailByID(ctx context.Context, id string) ([]*transaction_detail.TransactionDetail, error) {
	listTransactionDetail, err := uc.repoTransactionDetail.GetAllTransactionDetailByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return listTransactionDetail, nil
}

func (uc *TransactionDetailUseCaseInteractor) StoreTransactionDetail(ctx context.Context, trx_id int64, detail *transaction_detail.TransactionDetail) error {
	errStoreTransD := uc.repoTransactionDetail.StoreTransactionDetail(ctx, trx_id, detail)
	return errStoreTransD
}
