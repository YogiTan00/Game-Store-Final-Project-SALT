package transaction_test

import (
	"context"
	"game-store-final-project/project/internal/repository/mysql/transaction"
	"game-store-final-project/project/pkg/mysql_connection"
)

var (
	mysqlConn = mysql_connection.InitMysqlDB()
	//redisClient      = redis.InitRedisClient()
	ctx                        = context.Background()
	repoMysqlTransaction       = transaction.NewTransactionMysqlInteractor(mysqlConn)
	repoMysqlTransactionDetail = transaction.NewTransactionDetailMysqlInteractor(mysqlConn)
)
