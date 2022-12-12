package transaction

import (
	"errors"
	"game-store-final-project/project/domain/entity/transaction_detail"
	"time"
)

type Transaction struct {
	id               int
	customerId       int
	codeTransaction  string
	tanggalPembelian *time.Time
	transDetail      []*transaction_detail.TransactionDetail
	total            int64
}

type DTOTransaction struct {
	Id               int
	CustomerId       int
	CodeTransaction  string
	Tanggalpembelian *time.Time
	TransDetail      []*transaction_detail.DTOTransactionDetail
	Total            int64
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
	//date, _ := time.Parse("02-01-2006 15:04:05", t.Tanggalpembelian)

	return &Transaction{
		id:               t.Id,
		customerId:       t.CustomerId,
		codeTransaction:  t.CodeTransaction,
		tanggalPembelian: t.Tanggalpembelian,
		total:            t.Total,
	}, nil
}

func (t *Transaction) GetID() int {
	return t.id
}

func (t *Transaction) GetCustomerID() int {
	return t.customerId
}

func (t *Transaction) GetCodeTransaction() string {
	return t.codeTransaction
}

func (t *Transaction) GetTanggalPembelian() *time.Time {
	return t.tanggalPembelian
}

func (t *Transaction) GetTransDetail() []*transaction_detail.TransactionDetail {
	return t.transDetail
}

func (t *Transaction) GetTotal() int64 {
	return t.total
}

func (t *Transaction) AddTransDetail(dataTransDetail []*transaction_detail.TransactionDetail) *Transaction {
	t.transDetail = dataTransDetail
	return t
}
