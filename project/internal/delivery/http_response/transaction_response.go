package http_response

import (
	"encoding/json"
	"game-store-final-project/project/domain/entity/transaction"
)

type StatusTransaction struct {
	Code    int
	Mesaage string
}

type CustomReponseSingleTransaction struct {
	Status *StatusTransaction
	Data   *ResponseTransactionJson
}

type CustomReponseCollectionTransaction struct {
	Status *StatusTransaction
	Data   []*ResponseTransactionJson
}

type ResponseTransactionJson struct {
	Id              int    `json:"id"`
	CustomerId      int    `json:"customerId"`
	CodeTransaction string `json:"codeTransaction"`
}

func MapResponseListTransaction(dataTransaction []*transaction.Transaction, code int, message string) ([]byte, error) {
	listResponse := make([]*ResponseTransactionJson, 0)
	for _, data := range dataTransaction {
		response := &ResponseTransactionJson{
			Id:              data.GetID(),
			CustomerId:      data.GetCustomerID(),
			CodeTransaction: data.GetCodeTransaction(),
		}
		listResponse = append(listResponse, response)
	}

	httpResponse := &CustomReponseCollectionTransaction{
		Status: &StatusTransaction{
			Code:    code,
			Mesaage: message,
		},
		Data: listResponse,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
