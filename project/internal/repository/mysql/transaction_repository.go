package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"game-store-final-project/project/domain/entity"
	"game-store-final-project/project/domain/repository"
	"game-store-final-project/project/internal/repository/mysql/mapper"
	"game-store-final-project/project/internal/repository/mysql/model"
	"time"
)

type TransactionRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

func NewTransactionMysqlInteractor(conndb *sql.DB) repository.TransactionRepository {
	return &TransactionRepositoryMysqlInteractor{dbConn: conndb}
}

func (t TransactionRepositoryMysqlInteractor) GetAllTransaction(ctx context.Context) ([]*entity.Transaction, error) {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	stmt := fmt.Sprintf(`SELECT * FROM %s`, model.GetTableTransaction())
	rows, errMysql := t.dbConn.QueryContext(ctx, stmt)
	if errMysql != nil {
		return nil, errMysql
	}

	dataTransactionColletion := make([]*entity.Transaction, 0)
	for rows.Next() {
		var (
			id              int
			customerId      int
			codeTransaction string
		)
		err := rows.Scan(
			&id,
			&customerId,
			&codeTransaction,
		)
		if err != nil {
			return nil, err
		}
		trans, _ := time.Parse("INV02D01M2006Y15H04M05S", codeTransaction)
		dataTransaction, errTransaction := mapper.DataPTransDbToEntity(entity.DTOTransaction{
			Id:              id,
			CustomerId:      customerId,
			CodeTransaction: trans,
		})

		if errTransaction != nil {
			return nil, errTransaction
		}

		dataTransactionColletion = append(dataTransactionColletion, dataTransaction)

	}
	defer rows.Close()

	return dataTransactionColletion, nil
}
