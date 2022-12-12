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
	Id               int    `json:"id"`
	CustomerId       int    `json:"customerId"`
	CodeTransaction  string `json:"codeTransaction"`
	TanggalPembelian string `json:"tanggalPembelian"`
	Total            int64  `json:"total"`
}

func MapResponseTransaction(dataTransaction *transaction.Transaction, code int, message string) ([]byte, error) {
	response := &ResponseTransactionJson{
		Id:               dataTransaction.GetID(),
		CustomerId:       dataTransaction.GetCustomerID(),
		CodeTransaction:  dataTransaction.GetCodeTransaction(),
		TanggalPembelian: dataTransaction.GetTanggalPembelian().Format("02-01-2006 15:04:05"),
		Total:            dataTransaction.GetTotal(),
	}

	httpResponse := &CustomReponseSingleTransaction{
		Status: &StatusTransaction{
			Code:    code,
			Mesaage: message,
		},
		Data: response,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

func MapResponseListTransaction(dataTransaction []*transaction.Transaction, code int, message string) ([]byte, error) {
	listResponse := make([]*ResponseTransactionJson, 0)
	for _, data := range dataTransaction {
		response := &ResponseTransactionJson{
			Id:               data.GetID(),
			CustomerId:       data.GetCustomerID(),
			CodeTransaction:  data.GetCodeTransaction(),
			TanggalPembelian: data.GetTanggalPembelian().Format("02-01-2006 15:04:05"),
			Total:            data.GetTotal(),
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
