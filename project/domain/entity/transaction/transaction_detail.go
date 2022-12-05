package transaction

import (
	"errors"
	"game-store-final-project/project/domain/entity/item"
)

type TransactionDetail struct {
	id              int
	codeTransaction string
	transactionId   int
	itemId          string
	detailItem      *item.Item
	jumlahPembelian int
	hargaPembelian  int64
	total           int64
}

type DTOTransactionDetail struct {
	Id              int
	CodeTransaction string
	TransactionId   int
	ItemId          string
	DetailItem      *item.Item
	JumlahPembelian int
	HargaPembelian  int64
	Total           int64
}

func NewTransactionDetail(t DTOTransactionDetail) (*TransactionDetail, error) {
	if t.CodeTransaction == "" {
		return nil, errors.New("CODE TRANSACTION NOT SET")
	}
	if t.TransactionId == 0 {
		return nil, errors.New("TRANSACTION ID NOT SET")
	}
	if t.ItemId == "" || t.ItemId == "0" {
		return nil, errors.New("ITEM ID NOT SET")
	}

	return &TransactionDetail{
		id:              t.Id,
		codeTransaction: t.CodeTransaction,
		transactionId:   t.TransactionId,
		itemId:          t.ItemId,
		detailItem:      t.DetailItem,
		jumlahPembelian: t.JumlahPembelian,
		hargaPembelian:  t.HargaPembelian,
		total:           t.Total,
	}, nil
}

func (t *TransactionDetail) GetID() int {
	return t.id
}

func (t *TransactionDetail) GetCodeTransaction() string {
	return t.codeTransaction
}

func (t *TransactionDetail) GetTransactionID() int {
	return t.transactionId
}

func (t *TransactionDetail) GetItemID() string {
	return t.itemId
}

func (t *TransactionDetail) GetJumlahPembelian() int {
	return t.jumlahPembelian
}

func (t *TransactionDetail) GetHargaPembelian() int64 {
	return t.hargaPembelian
}

func (t *TransactionDetail) GetTotal() int64 {
	return t.total
}

func (t *TransactionDetail) GetDetail() *item.Item {
	return t.detailItem
}

func (t *TransactionDetail) AddDetail(data *item.Item) *TransactionDetail {
	t.detailItem = data
	return t
}
