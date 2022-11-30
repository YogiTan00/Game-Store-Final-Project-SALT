package mapper

import "game-store-final-project/project/domain/entity/transaction"

func DataTransactionDetailDbToEntity(dataDTO transaction.DTOTransactionDetail) (*transaction.TransactionDetail, error) {
	post, err := transaction.NewTransactionDetail(dataDTO)
	if err != nil {
		return nil, err
	}

	return post, nil
}
