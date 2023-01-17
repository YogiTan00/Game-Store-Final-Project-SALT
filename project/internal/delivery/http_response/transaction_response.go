package http_response

import (
	"encoding/json"
	"game-store-final-project/project/domain/entity/transaction"
	"time"
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
	Id               int                     `json:"id"`
	CustomerId       int                     `json:"customerId"`
	CodeTransaction  string                  `json:"codeTransaction"`
	TanggalPembelian *time.Time              `json:"tanggalPembelian"`
	Total            int64                   `json:"total"`
	TransDetail      []*ResponseTransactionD `json:"transDetail"`
}

type ResponseTransactionWithDetailJson struct {
	Id                int                              `json:"id"`
	CustomerId        int                              `json:"customerId"`
	CodeTransaction   string                           `json:"codeTransaction"`
	TanggalPembelian  time.Time                        `json:"tanggalPembelian"`
	Total             int64                            `json:"total"`
	TransactionDetail []*ResponseTransactionDetailJson `json:"transaction_detail"`
}

type ResponseTransactionD struct {
	Id            int                `json:"id"`
	TransactionId int                `json:"transactionId"`
	ItemId        int                `json:"itemId"`
	Detailitem    *ResponseItemDJson `json:"detailitem"`
	JumlahPembeli int                `json:"jumlahPembeli"`
	HargaPembeli  int64              `json:"hargaPembeli"`
	HargaDiscount int64              `json:"hargaDiscount"`
	Total         int64              `json:"total"`
}

type ResponseItemDJson struct {
	Nama     string `json:"nama"`
	Kategori string `json:"kategori"`
}

func MapResponseTransaction(dataTransaction *transaction.Transaction, code int, message string) ([]byte, error) {

	listTransD := make([]*ResponseTransactionD, 0)
	for _, data := range dataTransaction.GetTransDetail() {
		transD := &ResponseTransactionD{
			Id:            data.GetID(),
			TransactionId: data.GetTransactionID(),
			ItemId:        data.GetItemID(),
			Detailitem: &ResponseItemDJson{
				Nama:     data.GetDetail().GetNama(),
				Kategori: data.GetDetail().GetKategori(),
			},
			JumlahPembeli: data.GetJumlahPembelian(),
			HargaPembeli:  data.GetHargaPembelian(),
			HargaDiscount: data.GetHargaDiscount(),
			Total:         data.GetTotal(),
		}
		listTransD = append(listTransD, transD)
	}
	response := &ResponseTransactionJson{
		Id:               dataTransaction.GetID(),
		CustomerId:       dataTransaction.GetCustomerID(),
		CodeTransaction:  dataTransaction.GetCodeTransaction(),
		TanggalPembelian: dataTransaction.GetTanggalPembelian(),
		Total:            dataTransaction.GetTotal(),
		TransDetail:      listTransD,
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
		listTransD := make([]*ResponseTransactionD, 0)
		for _, dataTrans := range data.GetTransDetail() {
			transD := &ResponseTransactionD{
				Id:            dataTrans.GetID(),
				TransactionId: dataTrans.GetTransactionID(),
				ItemId:        dataTrans.GetItemID(),
				Detailitem: &ResponseItemDJson{
					Nama:     dataTrans.GetDetail().GetNama(),
					Kategori: dataTrans.GetDetail().GetKategori(),
				},
				JumlahPembeli: dataTrans.GetJumlahPembelian(),
				HargaPembeli:  dataTrans.GetHargaPembelian(),
				HargaDiscount: dataTrans.GetHargaDiscount(),
				Total:         dataTrans.GetTotal(),
			}
			listTransD = append(listTransD, transD)
		}

		response := &ResponseTransactionJson{
			Id:               data.GetID(),
			CustomerId:       data.GetCustomerID(),
			CodeTransaction:  data.GetCodeTransaction(),
			TanggalPembelian: data.GetTanggalPembelian(),
			Total:            data.GetTotal(),
			TransDetail:      listTransD,
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
