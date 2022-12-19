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

func (repo *VoucherRepositoryMysqlInteractor) GetVoucherByCode(ctx context.Context, code string) (*voucher.Voucher, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT * FROM voucher_customer where code_voucher = ? AND status = ? AND end_date > now()"
	rows, errMysql := repo.dbConn.QueryContext(ctx, sqlQuery, code, 1)
	if errMysql != nil {
		return nil, errMysql
	}

	if rows.Next() {
		var (
			id             int
			customer_id    int
			transaction_id int
			code_voucher   string
			nama_voucher   string
			start_date     time.Time
			end_date       time.Time
			use_date       time.Time
			status         int
			nilai_disc     int
			type_disc      int
		)

		err := rows.Scan(&id, &customer_id, &transaction_id, &code_voucher, &nama_voucher, &start_date, &end_date, &use_date, &status, &nilai_disc, &type_disc)
		if err != nil {
			return nil, err
		}

		voucher, errBuildVoucher := voucher.NewVoucherSingle(voucher.DTOVoucher{
			Id:            id,
			CustomerId:    customer_id,
			TransactionId: transaction_id,
			CodeVoucher:   code_voucher,
			NamaVoucher:   nama_voucher,
			StartDate:     start_date,
			EndDate:       end_date,
			UseDate:       use_date,
			Status:        status,
			NilaiDisc:     nilai_disc,
			TypeDisc:      type_disc,
		})
		if errBuildVoucher != nil {
			return nil, errBuildVoucher
		}

		return voucher, nil
	} else {
		return nil, nil
	}
}
