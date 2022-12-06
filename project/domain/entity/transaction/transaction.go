package transaction

import (
	"errors"
	"time"
)

type Transaction struct {
	id              int
	customerId      int
	codeTransaction string
}

type DTOTransaction struct {
	Id                int
	CustomerId        int
	CodeTransaction   time.Time
	TanggalPembelian  time.Time
	DetailTransaction []DTODetailTransaction
}

type DetailTransaction struct {
	item_id          int
	jumlah_pembelian int
}

type DTODetailTransaction struct {
	ItemId          int
	JumlahPembelian int
}

func NewTransaction(t DTOTransaction) (*Transaction, error) {
	if t.CustomerId == 0 {
		return nil, errors.New("ID COSTOMER NOT SET")
	}

	return &Transaction{
		id:              t.Id,
		customerId:      t.CustomerId,
		codeTransaction: t.CodeTransaction.Format("INV02D01M2006Y15H04M05S"),
	}, nil
}

func FenchDataTransactionFromDB(dataDTO DTOTransaction) *Transaction {
	return &Transaction{
		id:              dataDTO.Id,
		customerId:      dataDTO.CustomerId,
		codeTransaction: dataDTO.CodeTransaction.Format("INV02D01M2006Y15H04M05S"),
	}
}

func NewBuildTransaction(customer_id int, tanggal_pembelian time.Time, voucher string, d DTODetailTransaction) (*Transaction, error) {

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
