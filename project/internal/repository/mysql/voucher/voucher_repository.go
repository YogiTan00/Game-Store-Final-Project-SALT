package voucher

import (
	"context"
	"database/sql"
	"game-store-final-project/project/domain/entity/voucher"
	"game-store-final-project/project/domain/repository"
	"time"
)

// interaksi dengan DB
type VoucherRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

// build structnya, yang mengacu ke connection dan kontrak interface di repository
func NewVoucherRepositoryMysqlInteractor(connectionDatabse *sql.DB) repository.VoucherRepository {
	return &VoucherRepositoryMysqlInteractor{dbConn: connectionDatabse}
}

// implementasi dari interface kontrak dalam bentuk method receiver
func (repo *VoucherRepositoryMysqlInteractor) StoreVoucher(ctx context.Context, dataVoucher *voucher.Voucher) error {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// query insert to table voucher_customer
	insertQuery := "INSERT INTO voucher_customer(customer_id, code_voucher, nama_voucher, start_date, " +
		"end_date, use_date, status, nilai_disc, type_disc) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, errMysql = repo.dbConn.Exec(insertQuery, dataVoucher.GetCustomerId(), dataVoucher.GetCodeVoucher(), dataVoucher.GetNamaVoucher(),
		dataVoucher.GetStartDate(), dataVoucher.GetEndDate(), dataVoucher.GetUseDate(), dataVoucher.GetStatus(), dataVoucher.GetNilaiDisc(), dataVoucher.GetTypeDisc())

	if errMysql != nil {
		return errMysql
	}

	return nil
}
