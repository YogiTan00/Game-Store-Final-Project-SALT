package voucher_test

import (
	"context"
	"fmt"
	"game-store-final-project/project/internal/repository/mysql/voucher"
	"game-store-final-project/project/pkg/mysql_connection"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	mysqlConn                = mysql_connection.InitMysqlDB()
	ctx                      = context.Background()
	repoMysqlVoucherCustomer = voucher.NewVoucherRepositoryMysqlInteractor(mysqlConn)
)

func TestVoucherRepositoryMysqlInteractor_GetVoucherByCustomerId(t *testing.T) {
	voucherCustomer, err := repoMysqlVoucherCustomer.GetVoucherByCustomerId(ctx, "1")
	fmt.Println(voucherCustomer)
	assert.NotNil(t, voucherCustomer)
	assert.Nil(t, err)
}
