package transaction

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/domain/repository"
	"game-store-final-project/project/internal/repository/mysql/mapper"
	"game-store-final-project/project/internal/repository/mysql/model"
	"time"

	"github.com/rocketlaunchr/dbq/v2"
)

type TransactionRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

func NewTransactionMysqlInteractor(conndb *sql.DB) repository.TransactionRepository {
	return &TransactionRepositoryMysqlInteractor{dbConn: conndb}
}

func (t *TransactionRepositoryMysqlInteractor) GetTransactionByID(ctx context.Context, id string) (*transaction.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE id =?`, model.GetTableTransaction())
	opts := &dbq.Options{
		SingleResult:   true,
		ConcreteStruct: model.TransactionModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, t.dbConn, stmt, opts, id)
	if result == nil {
		return nil, errors.New("TRANSACTION NOT FOUND")
	}

	dataTransaction, errMap := mapper.ModelToDomainTransaction(result.(*model.TransactionModel))
	if errMap != nil {
		return nil, errMap
	}

	return dataTransaction, nil
}

func (t *TransactionRepositoryMysqlInteractor) GetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, model.GetTableTransaction())
	opts := &dbq.Options{
		SingleResult:   false,
		ConcreteStruct: model.TransactionModel{},
		DecoderConfig:  dbq.StdTimeConversionConfig(),
	}
	result := dbq.MustQ(ctx, t.dbConn, stmt, opts)
	if result == nil {
		return nil, errors.New("TRANSACTION NOT FOUND")
	}

	dataTransaction, errMap := mapper.ListModelToDomainTransaction(result.([]*model.TransactionModel))
	if errMap != nil {
		return nil, errMap
	}

	return dataTransaction, nil
}

func (t *TransactionRepositoryMysqlInteractor) GetAllTransactionByCustomerID(ctx context.Context, id string) ([]*transaction.Transaction, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s WHERE customer_id=?`, model.GetTableTransaction())
	rows, errMysql := t.dbConn.QueryContext(ctx, stmt, id)
	if errMysql != nil {
		return nil, errMysql
	}

	dataTransactionColletion := make([]*transaction.Transaction, 0)
	for rows.Next() {
		var (
			idTrans          int
			customerId       int
			codeTransaction  string
			tanggalPembelian *time.Time
			total            int64
		)
		err := rows.Scan(
			&idTrans,
			&customerId,
			&codeTransaction,
			&tanggalPembelian,
			&total,
		)
		if err != nil {
			return nil, err
		}

		dataTransaction, errTransaction := mapper.DataDbToEntityTransaction(transaction.DTOTransaction{
			Id:               idTrans,
			CustomerId:       customerId,
			CodeTransaction:  codeTransaction,
			Tanggalpembelian: tanggalPembelian,
			Total:            total,
		})

		if errTransaction != nil {
			return nil, errTransaction
		}

		dataTransactionColletion = append(dataTransactionColletion, dataTransaction)

	}
	defer rows.Close()

	return dataTransactionColletion, nil
}

func (t *TransactionRepositoryMysqlInteractor) StoreTransaction(ctx context.Context, dataTransaction *transaction.Transaction) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// query insert to table article
	insertQuery := "INSERT INTO transactions(customer_id, code_transaction, tanggal_pembelian, total) VALUES(?, ?, ?, ?)"

	res, errMysql := t.dbConn.Exec(insertQuery, dataTransaction.GetCustomerID(), dataTransaction.GetCodeTransaction(), dataTransaction.GetTanggalPembelian(),
		dataTransaction.GetTotal())

	if errMysql != nil {
		return 0, errMysql
	}

	// select last id
	lastInsertId, err := res.LastInsertId()
	if err != nil {
		return 0, errMysql
	}

	return lastInsertId, nil
}
