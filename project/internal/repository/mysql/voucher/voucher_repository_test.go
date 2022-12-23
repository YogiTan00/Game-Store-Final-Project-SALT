package voucher_test

import (
	"context"
	"fmt"
	entity_voucher "game-store-final-project/project/domain/entity/voucher"
	"game-store-final-project/project/internal/config/database/mysql"
	"game-store-final-project/project/internal/repository/mysql/voucher"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mysqlConn                = mysql.InitMysqlDB()
	ctx                      = context.Background()
	repoMysqlVoucherCustomer = voucher.NewVoucherRepositoryMysqlInteractor(mysqlConn)
)

func TestVoucherRepositoryMysqlInteractor_StoreVoucher(t *testing.T) {
	voucher, _ := entity_voucher.NewVoucherSingle(entity_voucher.DTOVoucher{
		CustomerId:    1,
		TransactionId: 1,
		CodeVoucher:   "PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M7",
		NamaVoucher:   "PREMI",
		StartDate:     time.Now(),
		EndDate:       time.Now(),
		UseDate:       time.Now(),
		Status:        1,
		NilaiDisc:     15,
		TypeDisc:      2,
	})

	err := repoMysqlVoucherCustomer.StoreVoucher(ctx, voucher)
	assert.Nil(t, err)
}

func TestVoucherRepositoryMysqlInteractor_GetVoucherByCode(t *testing.T) {
	voucherCustomer, err := repoMysqlVoucherCustomer.GetVoucherByCode(ctx, "PREMI-20D12M2022Y14H16M51SKFGQO4YLR1GWR")
	assert.NotNil(t, voucherCustomer)
	assert.Nil(t, err)
}

func TestVoucherRepositoryMysqlInteractor_GetVoucherByCustomerId(t *testing.T) {
	voucherCustomer, err := repoMysqlVoucherCustomer.GetVoucherByCustomerId(ctx, "1")
	fmt.Println(voucherCustomer)
	assert.NotNil(t, voucherCustomer)
	assert.Nil(t, err)
}

func TestVoucherRepositoryMysqlInteractor_UpdateVoucherById(t *testing.T) {
	err := repoMysqlVoucherCustomer.UpdateVoucherById(ctx, 102)
	assert.Nil(t, err)
}
