package transaction_detail

import (
	"errors"
	"game-store-final-project/project/domain/entity/item"
)

type TransactionDetail struct {
	id              int
	transactionId   int
	itemId          int
	detailItem      *item.Item
	jumlahPembelian int
	hargaPembelian  int64
	hargaDiscount   int64
	total           int64
}

type DTOTransactionDetail struct {
	Id              int
	TransactionId   int
	ItemId          int
	DetailItem      *item.Item
	JumlahPembelian int
	HargaPembelian  int64
	HargaDiscount   int64
	Total           int64
}

func NewTransactionDetail(t DTOTransactionDetail) (*TransactionDetail, error) {
	if t.TransactionId == 0 {
		return nil, errors.New("TRANSACTION ID NOT SET")
	}
	if t.ItemId == 0 {
		return nil, errors.New("ITEM ID NOT SET")
	}

	return &TransactionDetail{
		id:              t.Id,
		transactionId:   t.TransactionId,
		itemId:          t.ItemId,
		detailItem:      t.DetailItem,
		jumlahPembelian: t.JumlahPembelian,
		hargaPembelian:  t.HargaPembelian,
		hargaDiscount:   t.HargaDiscount,
		total:           t.Total,
	}, nil
}

func NewTransactionDetailWithoutTrxId(t DTOTransactionDetail) (*TransactionDetail, error) {
	if t.ItemId == 0 {
		return nil, errors.New("ITEM ID NOT SET")
	}

	return &TransactionDetail{
		id:              t.Id,
		transactionId:   t.TransactionId,
		itemId:          t.ItemId,
		detailItem:      t.DetailItem,
		jumlahPembelian: t.JumlahPembelian,
		hargaPembelian:  t.HargaPembelian,
		hargaDiscount:   t.HargaDiscount,
		total:           t.Total,
	}, nil
}

func (t *TransactionDetail) GetID() int {
	return t.id
}

func (t *TransactionDetail) GetTransactionID() int {
	return t.transactionId
}

func (t *TransactionDetail) GetItemID() int {
	return t.itemId
}

func (t *TransactionDetail) GetJumlahPembelian() int {
	return t.jumlahPembelian
}

func (t *TransactionDetail) GetHargaPembelian() int64 {
	return t.hargaPembelian
}

func (t *TransactionDetail) GetHargaDiscount() int64 {
	return t.hargaDiscount
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
