package voucher

import (
	"context"
	"database/sql"
	"game-store-final-project/project/domain/entity/customer"
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

func (repo *VoucherRepositoryMysqlInteractor) GetVoucherByCustomerId(ctx context.Context, id string) (*customer.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT * FROM customers where nik = ?"
	rows, errMysql := repo.dbConn.QueryContext(ctx, sqlQuery, id)
	if errMysql != nil {
		return nil, errMysql
	}

	if rows.Next() {
		var (
			id            int
			nik           string
			nama          string
			alamat        string
			no_tlp        string
			jenis_kelamin string
		)
		// first row
		err := rows.Scan(&id, &nik, &nama, &alamat, &no_tlp, &jenis_kelamin)
		if err != nil {
			return nil, err
		}

		// build struct customer
		customer, errBuildCustomer := customer.NewCustomer(customer.DTOCustomer{
			Id:           id,
			Nik:          nik,
			Nama:         nama,
			Alamat:       alamat,
			NoTlp:        no_tlp,
			JenisKelamin: jenis_kelamin,
		})
		if errBuildCustomer != nil {
			return nil, errBuildCustomer
		}
		sqlQueryVoucher := "SELECT * FROM voucher_customer where customer_id = ?"
		rowsVoucher, errVoucherMysql := repo.dbConn.QueryContext(ctx, sqlQueryVoucher, id)
		if errVoucherMysql != nil {
			return nil, errVoucherMysql
		}

		for rowsVoucher.Next() {
			var (
				customerId  int
				codeVoucher string
				namaVoucher string
				startDate   time.Time
				endDate     time.Time
				useDate     time.Time
				status      int
				nilaiDisc   int
				typeDisc    int
			)
			// first row
			err := rowsVoucher.Scan(&customerId, &codeVoucher, &namaVoucher, &startDate, &endDate, &useDate, status, nilaiDisc, typeDisc)
			if err != nil {
				return nil, err
			}

			// build struct customer
			voucher, errBuildVoucher := voucher.NewVoucher([]*voucher.DTOVoucher{
				{CustomerId: customerId,
					CodeVoucher: codeVoucher,
					NamaVoucher: namaVoucher,
					StartDate:   startDate,
					EndDate:     endDate,
					UseDate:     useDate,
					Status:      status,
					NilaiDisc:   nilaiDisc,
					TypeDisc:    typeDisc},
			})
			if errBuildVoucher != nil {
				return nil, errBuildVoucher
			}
			customer.AddVoucher(voucher)
		}
	}
	return nil, nil
}
