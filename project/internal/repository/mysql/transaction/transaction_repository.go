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

type TransactionRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

func NewTransactionMysqlInteractor(conndb *sql.DB) repository.TransactionRepository {
	return &TransactionRepositoryMysqlInteractor{dbConn: conndb}
}

func (t *TransactionRepositoryMysqlInteractor) GetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error) {
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

	dataTransactionColletion := make([]*transaction.Transaction, 0)
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
		dataTransaction, errTransaction := mapper.DataTransactionDbToEntity(transaction.DTOTransaction{
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
