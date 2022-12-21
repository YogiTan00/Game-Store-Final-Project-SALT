package transaction_detail

import (
	"context"
	"database/sql"
	"fmt"
	item2 "game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/domain/entity/transaction_detail"
	"game-store-final-project/project/domain/repository"
	"game-store-final-project/project/internal/repository/mysql/mapper"
	"game-store-final-project/project/internal/repository/mysql/model"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

type TransactionDetailRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

func NewTransactionDetailMysqlInteractor(conndb *sql.DB) repository.TransactionDetailRepository {
	return &TransactionDetailRepositoryMysqlInteractor{dbConn: conndb}
}

func (t *TransactionDetailRepositoryMysqlInteractor) GetAllTransactionDetail(ctx context.Context) ([]*transaction_detail.TransactionDetail, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	transactionDetail := model.GetTableTransactionDetail()
	item := model.GetTableItem()
	stmt := fmt.Sprintf(`SELECT t.id,t.transaction_id,t.item_id,i.nama,i.kategori,t.jumlah_pembelian,t.harga_pembelian,t.harga_discount,t.total FROM %s t JOIN %s i ON t.item_id = i.id`, transactionDetail, item)
	rows, errMysql := t.dbConn.QueryContext(ctx, stmt)
	if errMysql != nil {
		return nil, errMysql
	}
	dataPostCollection := make([]*transaction_detail.TransactionDetail, 0)
	for rows.Next() {
		var (
			idTransD        int
			transactionId   int
			itemId          int
			nama            string
			kategori        string
			jumlahPembelian int
			hargaPembelian  int64
			hargaDiscount   int64
			total           int64
		)
		err := rows.Scan(
			&idTransD,
			&transactionId,
			&itemId,
			&nama,
			&kategori,
			&jumlahPembelian,
			&hargaPembelian,
			&hargaDiscount,
			&total,
		)
		if err != nil {
			return nil, err
		}

		detailItem, errDetailitem := mapper.DataDbToEntityItem(item2.DTOItem{
			Nama:     nama,
			Kategori: kategori,
		})
		if errDetailitem != nil {
			return nil, errDetailitem
		}

		transDetail, errTransDetail := mapper.DataDbToEntityTransactionDetail(transaction_detail.DTOTransactionDetail{
			Id:              idTransD,
			TransactionId:   transactionId,
			ItemId:          itemId,
			DetailItem:      detailItem,
			JumlahPembelian: jumlahPembelian,
			HargaPembelian:  hargaPembelian,
			HargaDiscount:   hargaDiscount,
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

func (t *TransactionDetailRepositoryMysqlInteractor) GetTransactionDetailByID(ctx context.Context, id string) (*transaction_detail.TransactionDetail, error) {
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
		return nil, nil
	}

	listTransD, errMap := mapper.ModelToDomainTransactionDetail(result.(*model.TransactionDetailModel))
	if errMap != nil {
		return nil, errMap
	}

	return listTransD, nil
}

func (t *TransactionDetailRepositoryMysqlInteractor) GetAllTransactionDetailByID(ctx context.Context, id string) ([]*transaction_detail.TransactionDetail, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	transactionDetail := model.GetTableTransactionDetail()
	item := model.GetTableItem()
	stmt := fmt.Sprintf(`SELECT t.id,t.transaction_id,t.item_id,i.nama,i.kategori,t.jumlah_pembelian,t.harga_pembelian,t.harga_discount,t.total FROM %s t JOIN %s i ON t.item_id = i.id WHERE transaction_id =?`, transactionDetail, item)
	rows, errMysql := t.dbConn.QueryContext(ctx, stmt, id)
	if errMysql != nil {
		return nil, errMysql
	}
	dataPostCollection := make([]*transaction_detail.TransactionDetail, 0)
	for rows.Next() {
		var (
			idTransD        int
			transactionId   int
			itemId          int
			nama            string
			kategori        string
			jumlahPembelian int
			hargaPembelian  int64
			hargaDiscount   int64
			total           int64
		)
		err := rows.Scan(
			&idTransD,
			&transactionId,
			&itemId,
			&nama,
			&kategori,
			&jumlahPembelian,
			&hargaPembelian,
			&hargaDiscount,
			&total,
		)
		if err != nil {
			return nil, err
		}

		detailItem, errDetailitem := mapper.DataDbToEntityItem(item2.DTOItem{
			Nama:     nama,
			Kategori: kategori,
		})
		if errDetailitem != nil {
			return nil, errDetailitem
		}

		transDetail, errTransDetail := mapper.DataDbToEntityTransactionDetail(transaction_detail.DTOTransactionDetail{
			Id:              idTransD,
			TransactionId:   transactionId,
			ItemId:          itemId,
			DetailItem:      detailItem,
			JumlahPembelian: jumlahPembelian,
			HargaPembelian:  hargaPembelian,
			HargaDiscount:   hargaDiscount,
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

func (t *TransactionDetailRepositoryMysqlInteractor) StoreTransactionDetail(ctx context.Context, trx_id int64, detail *transaction_detail.TransactionDetail) error {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// query insert to table article
	insertQuery := "INSERT INTO transaction_detail(transaction_id, item_id, jumlah_pembelian, harga_pembelian, harga_discount, total) VALUES(?, ?, ?, ?, ?, ?)"

	_, errMysql = t.dbConn.Exec(insertQuery, trx_id, detail.GetItemID(), detail.GetJumlahPembelian(), detail.GetHargaPembelian(),
		detail.GetHargaDiscount(), detail.GetTotal())

	if errMysql != nil {
		return errMysql
	}

	return nil
}
