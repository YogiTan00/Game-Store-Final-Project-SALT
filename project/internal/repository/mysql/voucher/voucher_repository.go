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

func (repo *VoucherRepositoryMysqlInteractor) GetVoucherByCustomerId(ctx context.Context, id string) (*customer.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT * FROM customers where id = ?"
	rows, errMysql := repo.dbConn.QueryContext(ctx, sqlQuery, id)
	if errMysql != nil {
		return nil, errMysql
	}

	if rows.Next() {
		var (
			idCustomer    int
			nik           string
			nama          string
			alamat        string
			no_tlp        string
			jenis_kelamin string
		)
		// first row
		err := rows.Scan(&idCustomer, &nik, &nama, &alamat, &no_tlp, &jenis_kelamin)
		if err != nil {
			return nil, err
		}

		// build struct customer
		customer, errBuildCustomer := customer.NewCustomer(customer.DTOCustomer{
			Id:           idCustomer,
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
		listVoucher := make([]*voucher.Voucher, 0)
		for rowsVoucher.Next() {
			var (
				idVoucher      int
				customerId     int
				transaction_id int
				codeVoucher    string
				namaVoucher    string
				startDate      time.Time
				endDate        time.Time
				useDate        time.Time
				status         int
				nilaiDisc      int
				typeDisc       int
			)
			// first row
			errVoucher := rowsVoucher.Scan(&idVoucher, &customerId, &transaction_id, &codeVoucher, &namaVoucher,
				&startDate, &endDate, &useDate, &status, &nilaiDisc, &typeDisc)
			if errVoucher != nil {
				return nil, errVoucher
			}

			// build struct customer
			voucher, errBuildVoucher := voucher.NewVoucherCustomer(voucher.DTOVoucher{
				Id:            idVoucher,
				CustomerId:    customerId,
				TransactionId: transaction_id,
				CodeVoucher:   codeVoucher,
				NamaVoucher:   namaVoucher,
				StartDate:     startDate,
				EndDate:       endDate,
				UseDate:       useDate,
				Status:        status,
				NilaiDisc:     nilaiDisc,
				TypeDisc:      typeDisc,
			})
			if errBuildVoucher != nil {
				return nil, errBuildVoucher
			}
			listVoucher = append(listVoucher, voucher)
		}
		defer rowsVoucher.Close()
		customer.AddVoucher(listVoucher)
		return customer, nil
	} else {
		return nil, nil
	}
}

func (repo *VoucherRepositoryMysqlInteractor) UpdateVoucherById(ctx context.Context, id int) error {
	currentTime := time.Now()

	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// query insert to table voucher_customer
	insertQuery := "UPDATE voucher_customer SET use_date = ?, status = ? WHERE id = ?"

	_, errMysql = repo.dbConn.Exec(insertQuery, currentTime.Format("2006-01-02"), 0, id)

	if errMysql != nil {
		return errMysql
	}

	return nil
}
