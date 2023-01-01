package item

import (
	"context"
	"fmt"
	"game-store-final-project/project/internal/config/database/mysql"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mysqlConn                = mysql.InitMysqlDB()
	ctx                      = context.Background()
	repoMysqlTransactionItem = NewItemMysqlInteractor(mysqlConn)
)

func TestItemRepositoryMysqlInteractor_GetAllItem(t *testing.T) {
	transactionDetail, err := repoMysqlTransactionItem.GetAllItem(ctx)
	fmt.Println(transactionDetail)
	assert.NotNil(t, transactionDetail)
	assert.Nil(t, err)
}

func TestItemRepositoryMysqlInteractor_GetItemByID(t *testing.T) {
	transactionDetail, err := repoMysqlTransactionItem.GetItemByID(ctx, "1")
	fmt.Println(transactionDetail)
	assert.NotNil(t, transactionDetail)
	assert.Nil(t, err)
}
