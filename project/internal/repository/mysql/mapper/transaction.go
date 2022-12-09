package mapper

import (
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/internal/repository/mysql/model"
)

func DataDbToEntityTransaction(dataDTO transaction.DTOTransaction) (*transaction.Transaction, error) {
	post, err := transaction.NewTransaction(dataDTO)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func ModelToDomainTransaction(m *model.TransactionModel) (*transaction.Transaction, error) {
	transaction, err := transaction.NewTransaction(transaction.DTOTransaction{
		Id:               m.Id,
		CustomerId:       m.CustomerId,
		CodeTransaction:  m.CodeTransaction,
		Tanggalpembelian: m.TanggalPembelian,
		Total:            m.Total,
	})
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func ListModelToDomainTransaction(m []*model.TransactionModel) ([]*transaction.Transaction, error) {
	domains := make([]*transaction.Transaction, 0)
	for _, modelData := range m {
		d, _ := ModelToDomainTransaction(modelData)
		domains = append(domains, d)
	}
	return domains, nil
}
