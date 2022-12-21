package http_response

import (
	"encoding/json"
	"game-store-final-project/project/domain/entity/customer"
	"time"
)

type StatusVoucherCustomer struct {
	Code    int
	Mesaage string
}

type CustomReponseSingleVoucherCustomer struct {
	Status *StatusVoucherCustomer
	Data   *ResponseVoucherCustomerJson
}

type CustomReponseCollectionVoucherCustomer struct {
	Status *StatusVoucherCustomer
	Data   []*ResponseVoucherCustomerJson
}

type ResponseVoucherCustomerJson struct {
	Id           int                    `json:"id"`
	Nik          string                 `json:"nik"`
	Nama         string                 `json:"nama"`
	Alamat       string                 `json:"alamat"`
	NoTlp        string                 `json:"noTlp"`
	JenisKelamin string                 `json:"jenisKelamin"`
	Voucher      []*ResponseVoucherJson `json:"voucher"`
}

type ResponseVoucherJson struct {
	Id          int       `json:"id"`
	CustomerId  int       `json:"customerId"`
	CodeVoucher string    `json:"codeVoucher"`
	NamaVoucher string    `json:"namaVoucher"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	UseDate     time.Time `json:"useDate"`
	Status      int       `json:"status"`
	NilaiDisc   int       `json:"nilaiDisc"`
	TypeDisc    int       `json:"typeDisc"`
}

func MapResponseVoucherCustomer(dataVoucherCustomer *customer.Customer, code int, message string) ([]byte, error) {
	var resp *ResponseVoucherCustomerJson
	if dataVoucherCustomer != nil {
		listVoucher := make([]*ResponseVoucherJson, 0)
		for _, data := range dataVoucherCustomer.GetListVoucher() {
			dataVoucher := &ResponseVoucherJson{
				Id:          data.GetId(),
				CustomerId:  data.GetCustomerId(),
				CodeVoucher: data.GetCodeVoucher(),
				NamaVoucher: data.GetNamaVoucher(),
				StartDate:   data.GetStartDate(),
				EndDate:     data.GetEndDate(),
				UseDate:     data.GetUseDate(),
				Status:      data.GetStatus(),
				NilaiDisc:   data.GetNilaiDisc(),
				TypeDisc:    data.GetTypeDisc(),
			}
			listVoucher = append(listVoucher, dataVoucher)
		}
		resp = &ResponseVoucherCustomerJson{
			Id:           dataVoucherCustomer.GetId(),
			Nik:          dataVoucherCustomer.GetNik(),
			Nama:         dataVoucherCustomer.GetNama(),
			Alamat:       dataVoucherCustomer.GetAlamat(),
			NoTlp:        dataVoucherCustomer.GetNoTlp(),
			JenisKelamin: dataVoucherCustomer.GetJenisKelamin(),
			Voucher:      listVoucher,
		}
	}

	httpResponse := &CustomReponseSingleVoucherCustomer{
		Status: &StatusVoucherCustomer{
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
