package http_response

import (
	"encoding/json"
	"fmt"
	"game-store-final-project/project/domain/entity/customer"
)

type ReponseCustomer struct {
	Status *Status
	Data   *ReponseCustomerJson
}

type ReponseCustomerJson struct {
	Id           int                                  `json:"id"`
	Nik          string                               `json:"nik"`
	Nama         string                               `json:"nama"`
	Alamat       string                               `json:"alamat"`
	NoTlp        string                               `json:"no_tlp"`
	JenisKelamin string                               `json:"jenis_kelamin"`
	Transaction  []*ResponseTransactionWithDetailJson `json:"transaction"`
}

func MapResponseCustomer(code int, message string, customer *customer.Customer) ([]byte, error) {
	listRespTrx := make([]*ResponseTransactionWithDetailJson, 0)
	for _, trx := range customer.GetTrx() {
		listRespTrxDetail := make([]*ResponseTransactionDetailJson, 0)
		for _, trxDetail := range trx.GetDetailTrx() {
			fmt.Print(trxDetail.GetJumlahPembelian())
			respTrxDetail := &ResponseTransactionDetailJson{
				Id:            trxDetail.GetID(),
				TransactionId: trxDetail.GetTransactionID(),
				ItemId:        trxDetail.GetItemID(),
				JumlahPembeli: trxDetail.GetJumlahPembelian(),
				HargaPembeli:  trxDetail.GetHargaPembelian(),
				HargaDiscount: trxDetail.GetHargaDiscount(),
				Total:         trxDetail.GetTotal(),
			}

			listRespTrxDetail = append(listRespTrxDetail, respTrxDetail)
		}

		respTrx := &ResponseTransactionWithDetailJson{
			Id:                trx.GetID(),
			CustomerId:        trx.GetCustomerID(),
			CodeTransaction:   trx.GetCodeTransaction(),
			TanggalPembelian:  *trx.GetTanggalPembelian(),
			Total:             trx.GetTotal(),
			TransactionDetail: listRespTrxDetail,
		}

		listRespTrx = append(listRespTrx, respTrx)
	}

	listResponse := &ReponseCustomerJson{
		Id:           customer.GetId(),
		Nik:          customer.GetNik(),
		Nama:         customer.GetNama(),
		Alamat:       customer.GetAlamat(),
		NoTlp:        customer.GetNoTlp(),
		JenisKelamin: customer.GetJenisKelamin(),
		Transaction:  listRespTrx,
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
