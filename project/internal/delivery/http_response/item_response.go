package http_response

import (
	"encoding/json"
	item2 "game-store-final-project/project/domain/entity/item"
)

type StatusItem struct {
	Code    int
	Mesaage string
}

type CustomReponseSingleItem struct {
	Status *StatusItem
	Data   *ResponseItemJson
}

type CustomReponseCollectionItem struct {
	Status *StatusItem
	Data   []*ResponseItemJson
}

type ResponseItemJson struct {
	Id       int    `json:"id"`
	Nama     string `json:"nama"`
	Kategori string `json:"categoryId"`
	Harga    int64  `json:"harga"`
	Jumlah   int    `json:"jumlah"`
}

func MapResponseItem(dataItem *item2.Item, code int, message string) ([]byte, error) {
	response := &ResponseItemJson{
		Id:       dataItem.GetID(),
		Nama:     dataItem.GetNama(),
		Kategori: dataItem.GetKategori(),
		Harga:    dataItem.GetHarga(),
		Jumlah:   dataItem.GetJumlah(),
	}

	httpResponse := &CustomReponseSingleItem{
		Status: &StatusItem{
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
func MapResponseListItem(dataItem []*item2.Item, code int, message string) ([]byte, error) {
	listResponse := make([]*ResponseItemJson, 0)
	for _, data := range dataItem {
		response := &ResponseItemJson{
			Id:       data.GetID(),
			Nama:     data.GetNama(),
			Kategori: data.GetKategori(),
			Harga:    data.GetHarga(),
			Jumlah:   data.GetJumlah(),
		}
		listResponse = append(listResponse, response)
	}
	httpResponse := &CustomReponseCollectionItem{
		Status: &StatusItem{
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
