package mysql_test

import (
	"context"
	"fmt"
	"game-store-final-project/project/internal/repository/mysql"
	"game-store-final-project/project/pkg/mysql_connection"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mysqlConn = mysql_connection.InitMysqlDB()
	//redisClient      = redis.InitRedisClient()
	ctx                  = context.Background()
	repoMysqlTransaction = mysql.NewTransactionMysqlInteractor(mysqlConn)
)

func TestTransactionRepositoryMysqlInteractor_GetTransaction(t *testing.T) {
	transaction, err := repoMysqlTransaction.GetAllTransaction(ctx)
	fmt.Println(transaction)
	assert.NotNil(t, transaction)
	assert.Nil(t, err)
}
