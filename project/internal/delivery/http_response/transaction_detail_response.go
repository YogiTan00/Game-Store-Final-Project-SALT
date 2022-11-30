package http_response

import (
	"encoding/json"
	"game-store-final-project/project/domain/entity/transaction"
)

type StatusTransactionDetail struct {
	Code    int
	Mesaage string
}

type CustomReponseSingleTransactionDetail struct {
	Status *StatusTransactionDetail
	Data   *ResponseTransactionDetailJson
}

type CustomReponseCollectionTransactionDetail struct {
	Status *StatusTransactionDetail
	Data   []*ResponseTransactionDetailJson
}

type ResponseTransactionDetailJson struct {
	Id            int `json:"id"`
	TransactionId int `json:"transactionId"`
	ItemId        int `json:"itemId"`
}

func MapResponseListTransactionDetail(dataTransactionDetail []*transaction.TransactionDetail, code int, message string) ([]byte, error) {
	listResponse := make([]*ResponseTransactionDetailJson, 0)
	for _, data := range dataTransactionDetail {
		response := &ResponseTransactionDetailJson{
			Id:            data.GetID(),
			TransactionId: data.GetTransactionID(),
			ItemId:        data.GetItemID(),
		}
		listResponse = append(listResponse, response)
	}

	httpResponse := &CustomReponseCollectionTransactionDetail{
		Status: &StatusTransactionDetail{
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
