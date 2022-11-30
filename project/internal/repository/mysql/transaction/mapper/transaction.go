package mapper

import (
	"game-store-final-project/project/domain/entity/transaction"
)

func DataTransactionDbToEntity(dataDTO transaction.DTOTransaction) (*transaction.Transaction, error) {
	post, err := transaction.NewTransaction(dataDTO)
	if err != nil {
		return nil, err
	}

	return post, nil
}

//
//func ModelToDomainTransaction(m *model.TransactionModel) *transaction.Transaction {
//	trans, _ := time.Parse("INV02D01M2006Y15H04M05S", m.CodeTransaction)
//	dataTransaction := transaction.FenchDataTransactionFromDB(transaction.DTOTransaction{
//		Id:              m.Id,
//		CustomerId:      m.CustomerId,
//		CodeTransaction: trans,
//	})
//
//	return dataTransaction
//}
//
//func TransactionModelToDomain(m *model.TransactionModel) (*transaction.Transaction, error) {
//	codeTransaction, _ := time.Parse("INV02D01M2006Y15H04M05S", m.CodeTransaction)
//	transaction, err := transaction.NewTransaction(transaction.DTOTransaction{
//		Id:              m.Id,
//		CustomerId:      m.CustomerId,
//		CodeTransaction: codeTransaction,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	return transaction, nil
//}
//
//func TransactionCollectionModelToDomain(m []*model.TransactionModel) []*transaction.Transaction {
//	domains := make([]*transaction.Transaction, 0)
//	for _, modelData := range m {
//		d := ModelToDomainTransaction(modelData)
//		domains = append(domains, d)
//	}
//	return domains
//}
