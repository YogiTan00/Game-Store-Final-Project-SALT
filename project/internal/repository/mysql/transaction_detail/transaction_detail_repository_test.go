package transaction_detail_test

import (
	"context"
	"fmt"
	trx_detail_entity "game-store-final-project/project/domain/entity/transaction_detail"
	"game-store-final-project/project/internal/config/database/mysql"
	"game-store-final-project/project/internal/repository/mysql/transaction_detail"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mysqlConn                  = mysql.InitMysqlDB()
	ctx                        = context.Background()
	repoMysqlTransactionDetail = transaction_detail.NewTransactionDetailMysqlInteractor(mysqlConn)
)

func TestTransactionDetailRepositoryMysqlInteractor_GetAllTransaction(t *testing.T) {
	transactionDetail, err := repoMysqlTransactionDetail.GetAllTransactionDetail(ctx)
	fmt.Println(transactionDetail)
	assert.NotNil(t, transactionDetail)
	assert.Nil(t, err)
}

func TestTransactionDetailRepositoryMysqlInteractor_GetTransactionDetailByID(t *testing.T) {
	transactionDetail, err := repoMysqlTransactionDetail.GetTransactionDetailByID(ctx, "1")
	fmt.Println(transactionDetail)
	assert.NotNil(t, transactionDetail)
	assert.Nil(t, err)
}

func TestTransactionDetailRepositoryMysqlInteractor_GetAllTransactionDetailByID(t *testing.T) {
	transactionDetail, err := repoMysqlTransactionDetail.GetAllTransactionDetailByID(ctx, "1")
	fmt.Println(transactionDetail)
	assert.NotNil(t, transactionDetail)
	assert.Nil(t, err)
}

func TestTransactionRepositoryMysqlInteractor_StoreTransactionDetail(t *testing.T) {
	transaction_detail, _ := trx_detail_entity.NewTransactionDetail(trx_detail_entity.DTOTransactionDetail{
		Id:              1,
		TransactionId:   1,
		ItemId:          1,
		JumlahPembelian: 1,
		HargaPembelian:  4360000,
		HargaDiscount:   0,
		Total:           4360000,
	})

	err := repoMysqlTransactionDetail.StoreTransactionDetail(ctx, 2, transaction_detail)
	assert.Nil(t, err)
}
