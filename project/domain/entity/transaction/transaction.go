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
	tanggalPembelian  time.Time
	total             int64
	hargaDiscount     int64
	totalHarga        int64
	itemPembelian     []*ItemPembelian
}

type DTOTransaction struct {
	Id                int
	VoucherCustomerId int
	CustomerId        int
	CodeTransaction   string
	Tanggalpembelian  string
	Total             int64
	HargaDiscount     int64
	TotalHarga        int64
	ItemPembelian     []*DTOItemPembelian
}

type ItemPembelian struct {
	itemId          int
	jumlahPembelian int
}

type DTOItemPembelian struct {
	ItemId          int
	JumlahPembelian int
}

func NewTransaction(t DTOTransaction) (*Transaction, error) {
	if t.CustomerId == 0 {
		return nil, errors.New("ID COSTOMER NOT SET")
	}

	// convert string to time
	date, _ := time.Parse("02-01-2006 15:04:05", t.Tanggalpembelian)

	items := make([]*ItemPembelian, 0)
	for _, item := range t.ItemPembelian {
		dataItem := &ItemPembelian{
			itemId:          item.ItemId,
			jumlahPembelian: item.JumlahPembelian,
		}
		items = append(items, dataItem)
	}

	return &Transaction{
		id:                t.Id,
		voucherCustomerId: t.VoucherCustomerId,
		customerId:        t.CustomerId,
		codeTransaction:   t.CodeTransaction,
		tanggalPembelian:  date,
		total:             t.Total,
		hargaDiscount:     t.HargaDiscount,
		totalHarga:        t.TotalHarga,
		itemPembelian:     items,
	}, nil
}

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

func (t *Transaction) GetTanggalPembelian() time.Time {
	return t.tanggalPembelian
}
