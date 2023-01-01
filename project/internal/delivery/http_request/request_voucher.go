package http_request

import "time"

type RequestVoucher struct {
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
