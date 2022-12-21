package test_data

import (
	"game-store-final-project/project/domain/entity/customer"
	voucher2 "game-store-final-project/project/domain/entity/voucher"
	"time"
)

func GetTestDataVoucherCustomer(count int) *customer.Customer {
	listVoucher := make([]*voucher2.Voucher, 0)
	for i := 1; i <= count; i++ {
		voucher, _ := voucher2.NewVoucherCustomer(voucher2.DTOVoucher{
			Id:          1,
			CustomerId:  1,
			CodeVoucher: "P",
			NamaVoucher: "P",
			StartDate:   time.Now(),
			EndDate:     time.Now().AddDate(1, 0, 0),
			UseDate:     time.Time{},
			NilaiDisc:   0,
			TypeDisc:    0,
		})
		listVoucher = append(listVoucher, voucher)
	}

	customer, _ := customer.NewCustomer(customer.DTOCustomer{
		Id:           1,
		Nik:          "1234567890123456",
		Nama:         "r",
		Alamat:       "q",
		NoTlp:        "123456789012",
		JenisKelamin: "LK",
		Voucher:      listVoucher,
	})
	return customer
}
