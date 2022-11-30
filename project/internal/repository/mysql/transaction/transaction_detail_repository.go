package transaction

import (
	"context"
	"database/sql"
	"fmt"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/domain/repository"
	"game-store-final-project/project/internal/repository/mysql/transaction/mapper"
	"game-store-final-project/project/internal/repository/mysql/transaction/model"
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
			id            int
			transactionID int
			itemID        int
		)
		err := rows.Scan(
			&id,
			&transactionID,
			&itemID,
		)
		if err != nil {
			return nil, err
		}
		dataTransactionDetail, errTransactionDetail := mapper.DataTransactionDetailDbToEntity(transaction.DTOTransactionDetail{
			Id:            id,
			TransactionId: transactionID,
			ItemId:        itemID,
		})

		if errTransactionDetail != nil {
			return nil, errTransactionDetail
		}

		dataTransactionColletionDetail = append(dataTransactionColletionDetail, dataTransactionDetail)

	}
	defer rows.Close()

	return dataTransactionColletionDetail, nil
}
