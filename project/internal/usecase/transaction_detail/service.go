package transaction_detail

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction_detail"
)

func (uc *TransactionDetailUseCaseInteractor) UcGetAllTransactionDetail(ctx context.Context) ([]*transaction_detail.TransactionDetail, error) {
	listTransactionDetail, err := uc.repoTransactionDetail.GetAllTransactionDetail(ctx)
	if err != nil {
		return nil, err
	}

	return listTransactionDetail, nil
}

func (uc *TransactionDetailUseCaseInteractor) UcGetTransactionDetailByID(ctx context.Context, id string) (*transaction_detail.TransactionDetail, error) {
	listTransactionDetail, err := uc.repoTransactionDetail.GetTransactionDetailByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return listTransactionDetail, nil
}

func (uc *TransactionDetailUseCaseInteractor) UcGetAllTransactionDetailByID(ctx context.Context, id string) ([]*transaction_detail.TransactionDetail, error) {
	listTransactionDetail, err := uc.repoTransactionDetail.GetAllTransactionDetailByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return listTransactionDetail, nil
}
