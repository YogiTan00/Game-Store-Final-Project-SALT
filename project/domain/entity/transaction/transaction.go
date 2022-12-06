package transaction

import (
	"errors"
	"time"
)

type Transaction struct {
	id                int
	voucherCustomerId int
	customerId        int
	codeTransaction   string
	tanggalPembelian  string
	total             int64
	hargaDiscount     int64
	totalHarga        int64
}

type DTOTransaction struct {
	Id                int
	VoucherCustomerId int
	CustomerId        int
	CodeTransaction   time.Time
	Tanggalpembelian  time.Time
	Total             int64
	HargaDiscount     int64
	TotalHarga        int64
}

func NewTransaction(t DTOTransaction) (*Transaction, error) {
	if t.CustomerId == 0 {
		return nil, errors.New("ID COSTOMER NOT SET")
	}

	return &Transaction{
		id:                t.Id,
		voucherCustomerId: t.VoucherCustomerId,
		customerId:        t.CustomerId,
		codeTransaction:   t.CodeTransaction.Format("INV02D01M2006Y15H04M05S"),
		tanggalPembelian:  t.Tanggalpembelian.Format("02-01-2006 15:04:05"),
		total:             t.Total,
		hargaDiscount:     t.HargaDiscount,
		totalHarga:        t.TotalHarga,
	}, nil
}

//func FenchDataTransactionFromDB(dataDTO DTOTransaction) *Transaction {
//	return &Transaction{
//		id:              dataDTO.Id,
//		customerId:      dataDTO.CustomerId,
//		codeTransaction: dataDTO.CodeTransaction.Format("INV02D01M2006Y15H04M05S"),
//	}
//}

// func NewBuildTransaction(customer_id int, tanggal_pembelian time.Time, voucher string, d DTODetailTransaction) (*Transaction, error) {

// }

func (t *Transaction) GetID() int {
	return t.id
}

func (t *Transaction) GetVoucherCustomerID() int {
	return t.voucherCustomerId
}

func (t *Transaction) GetCustomerID() int {
	return t.customerId
}

func (t *Transaction) GetCodeTransaction() string {
	return t.codeTransaction
}

func (t *Transaction) GetTanggalPembelian() string {
	return t.tanggalPembelian
}
