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
	Id            int    `json:"id"`
	CodeTrans     string `json:"codeTrans"`
	TransactionId int    `json:"transactionId"`
	ItemId        string `json:"itemId"`
	JumlahPembeli int    `json:"jumlahPembeli"`
	HargaPembeli  int64  `json:"hargaPembeli"`
	Total         int64  `json:"total"`
}

func MapResponseTransactionDetail(dataTransactionDetail *transaction.TransactionDetail, code int, message string) ([]byte, error) {
	var resp *ResponseTransactionDetailJson
	if dataTransactionDetail != nil {
		resp = &ResponseTransactionDetailJson{
			Id:            dataTransactionDetail.GetID(),
			TransactionId: dataTransactionDetail.GetTransactionID(),
			ItemId:        dataTransactionDetail.GetItemID(),
		}
	}

	httpResponse := &CustomReponseSingleTransactionDetail{
		Status: &StatusTransactionDetail{
			Code:    code,
			Mesaage: message,
		},
		Data: resp,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

func MapResponseListTransactionDetail(dataTransactionDetail []*transaction.TransactionDetail, code int, message string) ([]byte, error) {
	listResponse := make([]*ResponseTransactionDetailJson, 0)
	for _, data := range dataTransactionDetail {
		response := &ResponseTransactionDetailJson{
			Id:            data.GetID(),
			CodeTrans:     data.GetCodeTransaction(),
			TransactionId: data.GetTransactionID(),
			ItemId:        data.GetItemID(),
			JumlahPembeli: data.GetJumlahPembelian(),
			HargaPembeli:  data.GetHargaPembelian(),
			Total:         data.GetTotal(),
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
