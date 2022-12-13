package http_response

import (
	"encoding/json"
	"game-store-final-project/project/domain/entity/transaction_detail"
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
	Id            int                     `json:"id"`
	TransactionId int                     `json:"transactionId"`
	ItemId        int                     `json:"itemId"`
	Detailitem    *ResponseItemDetailJson `json:"detailitem"`
	JumlahPembeli int                     `json:"jumlahPembeli"`
	HargaPembeli  int64                   `json:"hargaPembeli"`
	HargaDiscount int64                   `json:"hargaDiscount"`
	Total         int64                   `json:"total"`
}

type ResponseItemDetailJson struct {
	Nama     string `json:"nama"`
	Kategori string `json:"kategori"`
}

func MapResponseTransactionDetail(dataTransactionDetail *transaction_detail.TransactionDetail, code int, message string) ([]byte, error) {
	var resp *ResponseTransactionDetailJson
	if dataTransactionDetail != nil {
		dataItem := &ResponseItemDetailJson{
			Nama:     dataTransactionDetail.GetDetail().GetNama(),
			Kategori: dataTransactionDetail.GetDetail().GetKategori(),
		}
		resp = &ResponseTransactionDetailJson{
			Id:            dataTransactionDetail.GetID(),
			TransactionId: dataTransactionDetail.GetTransactionID(),
			ItemId:        dataTransactionDetail.GetItemID(),
			Detailitem:    dataItem,
			JumlahPembeli: dataTransactionDetail.GetJumlahPembelian(),
			HargaPembeli:  dataTransactionDetail.GetHargaPembelian(),
			HargaDiscount: dataTransactionDetail.GetHargaDiscount(),
			Total:         dataTransactionDetail.GetTotal(),
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

func MapResponseListTransactionDetail(dataTransactionDetail []*transaction_detail.TransactionDetail, code int, message string) ([]byte, error) {
	listResponse := make([]*ResponseTransactionDetailJson, 0)
	for _, data := range dataTransactionDetail {
		dataItem := &ResponseItemDetailJson{
			Nama:     data.GetDetail().GetNama(),
			Kategori: data.GetDetail().GetKategori(),
		}
		response := &ResponseTransactionDetailJson{
			Id:            data.GetID(),
			TransactionId: data.GetTransactionID(),
			ItemId:        data.GetItemID(),
			Detailitem:    dataItem,
			JumlahPembeli: data.GetJumlahPembelian(),
			HargaPembeli:  data.GetHargaPembelian(),
			HargaDiscount: data.GetHargaDiscount(),
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
