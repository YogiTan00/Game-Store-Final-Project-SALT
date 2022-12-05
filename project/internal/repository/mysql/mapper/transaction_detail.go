package mapper

import (
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/internal/repository/mysql/model"
)

func DataTransactionDetailDbToEntity(dataDTO transaction.DTOTransactionDetail) (*transaction.TransactionDetail, error) {
	post, err := transaction.NewTransactionDetail(dataDTO)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func TransactionDetailModelToDomain(m *model.TransactionDetailModel) (*transaction.TransactionDetail, error) {

	transactionDetail, err := transaction.NewTransactionDetail(transaction.DTOTransactionDetail{
		Id:              m.Id,
		CodeTransaction: m.CodeTrans,
		TransactionId:   m.TransactionId,
		ItemId:          m.ItemId,
		JumlahPembelian: m.JumlahPembelian,
		HargaPembelian:  m.HargaPembelian,
		Total:           m.Total,
	})
	if err != nil {
		return nil, err
	}

	return transactionDetail, nil
}

func TransactionDetailListModelToDomain(m []*model.TransactionDetailModel) ([]*transaction.TransactionDetail, error) {
	domains := make([]*transaction.TransactionDetail, 0)
	for _, modelData := range m {
		d, _ := TransactionDetailModelToDomain(modelData)
		domains = append(domains, d)
	}
	return domains, nil
}
