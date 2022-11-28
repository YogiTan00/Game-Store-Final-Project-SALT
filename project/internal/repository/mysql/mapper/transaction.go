package mapper

import (
	"game-store-final-project/project/domain/entity"
	"game-store-final-project/project/internal/repository/mysql/model"
	"time"
)

func DataPTransDbToEntity(dataDTO entity.DTOTransaction) (*entity.Transaction, error) {
	post, err := entity.NewTransaction(dataDTO)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func ModelToDomainTransaction(m *model.TransactionModel) *entity.Transaction {
	transaction, _ := time.Parse("INV02D01M2006Y15H04M05S", m.CodeTransaction)
	dataTransaction := entity.FenchDataTransactionFromDB(entity.DTOTransaction{
		Id:              m.Id,
		CustomerId:      m.CustomerId,
		CodeTransaction: transaction,
	})

	return dataTransaction
}

func TransactionModelToDomain(m *model.TransactionModel) (*entity.Transaction, error) {
	codeTransaction, _ := time.Parse("INV02D01M2006Y15H04M05S", m.CodeTransaction)
	transaction, err := entity.NewTransaction(entity.DTOTransaction{
		Id:              m.Id,
		CustomerId:      m.CustomerId,
		CodeTransaction: codeTransaction,
	})
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func TransactionCollectionModelToDomain(m []*model.TransactionModel) []*entity.Transaction {
	domains := make([]*entity.Transaction, 0)
	for _, modelData := range m {
		d := ModelToDomainTransaction(modelData)
		domains = append(domains, d)
	}
	return domains
}
