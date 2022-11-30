package transaction

import "errors"

type TransactionDetail struct {
	id            int
	transactionId int
	itemId        int
}

type DTOTransactionDetail struct {
	Id            int
	TransactionId int
	ItemId        int
}

func NewTransactionDetail(t DTOTransactionDetail) (*TransactionDetail, error) {
	if t.TransactionId == 0 {
		return nil, errors.New("ID TRANSACTION NOT SET")
	}
	if t.ItemId == 0 {
		return nil, errors.New("ID ITEM NOT SET")
	}
	return &TransactionDetail{
		id:            t.Id,
		transactionId: t.TransactionId,
		itemId:        t.ItemId,
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
