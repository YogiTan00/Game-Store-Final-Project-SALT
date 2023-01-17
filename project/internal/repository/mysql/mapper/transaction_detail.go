package mapper

import (
	"game-store-final-project/project/domain/entity/transaction_detail"
	"game-store-final-project/project/internal/repository/mysql/model"
)

func DataDbToEntityTransactionDetail(dataDTO transaction_detail.DTOTransactionDetail) (*transaction_detail.TransactionDetail, error) {
	post, err := transaction_detail.NewTransactionDetail(dataDTO)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func ModelToDomainTransactionDetail(m *model.TransactionDetailModel) (*transaction_detail.TransactionDetail, error) {

	transactionDetail, err := transaction_detail.NewTransactionDetail(transaction_detail.DTOTransactionDetail{
		Id:              m.Id,
		TransactionId:   m.TransactionId,
		ItemId:          m.ItemId,
		JumlahPembelian: m.JumlahPembelian,
		HargaPembelian:  m.HargaPembelian,
		HargaDiscount:   m.HargaDiscount,
		Total:           m.Total,
	})
	if err != nil {
		return nil, err
	}

	return transactionDetail, nil
}

func ListModelToDomainTransactionDetail(m []*model.TransactionDetailModel) ([]*transaction_detail.TransactionDetail, error) {
	domains := make([]*transaction_detail.TransactionDetail, 0)
	for _, modelData := range m {
		d, _ := ModelToDomainTransactionDetail(modelData)
		domains = append(domains, d)
	}
	return domains, nil
}
