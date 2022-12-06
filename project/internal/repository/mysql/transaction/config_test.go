package transaction_test

import (
	"context"
	item2 "game-store-final-project/project/internal/repository/mysql/item"
	"game-store-final-project/project/internal/repository/mysql/transaction"
	"game-store-final-project/project/pkg/mysql_connection"
)

var (
	mysqlConn = mysql_connection.InitMysqlDB()
	//redisClient      = redis.InitRedisClient()
	ctx                        = context.Background()
	repoMysqlTransaction       = transaction.NewTransactionMysqlInteractor(mysqlConn)
	repoMysqlTransactionDetail = transaction.NewTransactionDetailMysqlInteractor(mysqlConn)
	RepoMysqlItem              = item2.NewItemMysqlInteractor(mysqlConn)
)
