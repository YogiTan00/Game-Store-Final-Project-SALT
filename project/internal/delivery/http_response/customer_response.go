package http_response

import (
	"encoding/json"
	"fmt"
	"game-store-final-project/project/domain/entity/customer"
	"game-store-final-project/project/domain/entity/transaction"
)

type ReponseCustomer struct {
	Status *Status
	Data   *ReponseCustomerJson
}

type ReponseCustomerJson struct {
	Id           int                        `json:"id"`
	Nik          string                     `json:"nik"`
	Nama         string                     `json:"nama"`
	Alamat       string                     `json:"alamat"`
	NoTlp        string                     `json:"no_tlp"`
	JenisKelamin string                     `json:"jenis_kelamin"`
	Transaction  []*transaction.Transaction `json:"transaction"`
}

func MapResponseCustomer(code int, message string, customer *customer.Customer) ([]byte, error) {
	listResponse := &ReponseCustomerJson{
		Id:           customer.GetId(),
		Nik:          customer.GetNik(),
		Nama:         customer.GetNama(),
		Alamat:       customer.GetAlamat(),
		NoTlp:        customer.GetNoTlp(),
		JenisKelamin: customer.GetJenisKelamin(),
		Transaction:  customer.GetDetailTrx(),
	}

	for _, trx := range customer.GetDetailTrx() {
		fmt.Print(trx)
	}

	httpResponse := &ReponseCustomer{
		Status: &Status{
			Code:    code,
			Message: message,
		},
		Data: listResponse,
	}

	respJson, err := json.Marshal(httpResponse)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}
