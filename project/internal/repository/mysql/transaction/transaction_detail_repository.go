package transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	item2 "game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/domain/repository"
	"game-store-final-project/project/internal/repository/mysql/mapper"
	"game-store-final-project/project/internal/repository/mysql/model"
	"github.com/rocketlaunchr/dbq/v2"
	"time"
)

type TransactionDetailRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

func NewTransactionDetailMysqlInteractor(conndb *sql.DB) repository.TransactionDetailRepository {
	return &TransactionDetailRepositoryMysqlInteractor{dbConn: conndb}
}

func (t *TransactionDetailRepositoryMysqlInteractor) GetAllTransactionDetail(ctx context.Context) ([]*transaction.TransactionDetail, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, model.GetTableTransactionDetail())
	rows, errMysql := t.dbConn.QueryContext(ctx, stmt)
	if errMysql != nil {
		return nil, errMysql
	}

	dataTransactionColletionDetail := make([]*transaction.TransactionDetail, 0)
	for rows.Next() {
		var (
			id              int
			codeTrans       string
			transactionID   int
			itemID          string
			jumlahPembelian int
			hargaPembelian  int64
			total           int64
		)
		err := rows.Scan(
			&id,
			&codeTrans,
			&transactionID,
			&itemID,
			&jumlahPembelian,
			&hargaPembelian,
			&total,
		)
		if err != nil {
			return nil, err
		}
		dataTransactionDetail, errTransactionDetail := mapper.DataTransactionDetailDbToEntity(transaction.DTOTransactionDetail{
			Id:              id,
			CodeTransaction: codeTrans,
			TransactionId:   transactionID,
			ItemId:          itemID,
			JumlahPembelian: jumlahPembelian,
			HargaPembelian:  hargaPembelian,
			Total:           total,
		})

		if errTransactionDetail != nil {
			return nil, errTransactionDetail
		}

		dataTransactionColletionDetail = append(dataTransactionColletionDetail, dataTransactionDetail)

	}
	defer rows.Close()

	return dataTransactionColletionDetail, nil
}

func (t *TransactionDetailRepositoryMysqlInteractor) GetTransactionDetailByID(ctx context.Context, id string) (*transaction.TransactionDetail, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	transactionDetail := model.GetTableTransactionDetail()
	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE transaction_id =?`, transactionDetail)
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: model.TransactionDetailModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, t.dbConn, stmt, opts, id)
	if result == nil {
		return nil, errors.New("ITEM NOT FOUND")
	}

	listTransD, errMap := mapper.TransactionDetailModelToDomain(result.(*model.TransactionDetailModel))
	if errMap != nil {
		return nil, errMap
	}

	return listTransD, nil
}

func (t *TransactionDetailRepositoryMysqlInteractor) GetAllTransactionDetailByID(ctx context.Context, id string) ([]*transaction.TransactionDetail, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	transactionDetail := model.GetTableTransactionDetail()
	item := model.GetTableItem()
	stmt := fmt.Sprintf(`SELECT t.id,t.code_transaction,t.transaction_id,t.item_id,i.nama,i.kategori,t.jumlah_pembelian,t.harga_pembelian,t.total FROM %s t JOIN %s i ON t.item_id = i.id WHERE transaction_id =?`, transactionDetail, item)
	rows, errMysql := t.dbConn.QueryContext(ctx, stmt, id)
	if errMysql != nil {
		return nil, errMysql
	}
	dataPostCollection := make([]*transaction.TransactionDetail, 0)
	for rows.Next() {
		var (
			idTransD        int
			codeTrans       string
			transactionId   int
			itemId          string
			nama            string
			kategori        string
			jumlahPembelian int
			hargaPembelian  int64
			total           int64
		)
		err := rows.Scan(
			&idTransD,
			&codeTrans,
			&transactionId,
			&itemId,
			&nama,
			&kategori,
			&jumlahPembelian,
			&hargaPembelian,
			&total,
		)
		if err != nil {
			return nil, err
		}

		detailItem, errDetailitem := mapper.DataItemDbToEntity(item2.DTOItem{
			Nama:     nama,
			Kategori: kategori,
		})
		if errDetailitem != nil {
			return nil, errDetailitem
		}

		transDetail, errTransDetail := mapper.DataTransactionDetailDbToEntity(transaction.DTOTransactionDetail{
			Id:              idTransD,
			CodeTransaction: codeTrans,
			TransactionId:   transactionId,
			ItemId:          itemId,
			DetailItem:      detailItem,
			JumlahPembelian: jumlahPembelian,
			HargaPembelian:  hargaPembelian,
			Total:           total,
		})
		if errTransDetail != nil {
			return nil, errTransDetail
		}
		dataPostCollection = append(dataPostCollection, transDetail)
	}
	defer rows.Close()

	return dataPostCollection, nil
}
