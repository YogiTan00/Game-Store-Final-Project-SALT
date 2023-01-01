package customer

import (
	"context"
	"database/sql"
	"game-store-final-project/project/domain/entity/customer"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/domain/entity/transaction_detail"
	"game-store-final-project/project/domain/repository"
	"time"
)

// interaksi dengan DB
type CustomerRepositoryMysqlInteractor struct {
	dbConn *sql.DB
}

// build structnya, yang mengacu ke connection dan kontrak interface di repository
func NewCustomerRepositoryMysqlInteractor(connectionDatabse *sql.DB) repository.CustomerRepository {
	return &CustomerRepositoryMysqlInteractor{dbConn: connectionDatabse}
}

// implementasi dari interface kontrak dalam bentuk method receiver
func (repo *CustomerRepositoryMysqlInteractor) StoreCustomer(ctx context.Context, dataCustomer *customer.Customer) error {
	var (
		errMysql error
	)

	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// query insert to table article
	insertQuery := "INSERT INTO customers(nik, nama, alamat, no_tlp, " +
		"jenis_kelamin) VALUES(?, ?, ?, ?, ?)"

	_, errMysql = repo.dbConn.Exec(insertQuery, dataCustomer.GetNik(), dataCustomer.GetNama(), dataCustomer.GetAlamat(),
		dataCustomer.GetNoTlp(), dataCustomer.GetJenisKelamin())

	if errMysql != nil {
		return errMysql
	}

	return nil
}

func (repo *CustomerRepositoryMysqlInteractor) IndexCustomerWithTransaction(ctx context.Context, nik string) (*customer.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT * FROM customers where nik = ?"
	rows, errMysql := repo.dbConn.QueryContext(ctx, sqlQuery, nik)
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

		sqlQueryTrx := "SELECT * FROM transactions where customer_id = ?"
		rowsTrx, errTrxMysql := repo.dbConn.QueryContext(ctx, sqlQueryTrx, id)
		if errTrxMysql != nil {
			return nil, errTrxMysql
		}

		// slice trx
		trxMany := make([]*transaction.Transaction, 0)
		for rowsTrx.Next() {
			var (
				idTrx             int
				customer_id       int
				code_transaction  string
				tanggal_pembelian time.Time
				total             int64
			)

			err := rowsTrx.Scan(&idTrx, &customer_id, &code_transaction, &tanggal_pembelian, &total)
			if err != nil {
				return nil, err
			}

			sqlQueryTrxDetail := "SELECT * FROM transaction_detail where transaction_id = ?"
			rowsTrxDetail, errTrxMysqlDetail := repo.dbConn.QueryContext(ctx, sqlQueryTrxDetail, idTrx)
			if errTrxMysqlDetail != nil {
				return nil, errTrxMysqlDetail
			}

			// slice detail trx
			trxDetailMany := make([]*transaction_detail.TransactionDetail, 0)
			for rowsTrxDetail.Next() {
				var (
					idTrxDetail      int
					transaction_id   int
					item_id          int
					jumlah_pembelian int
					harga_pembelian  int64
					harga_discount   int64
					total            int64
				)

				err := rowsTrxDetail.Scan(&idTrxDetail, &transaction_id, &item_id, &jumlah_pembelian, &harga_pembelian, &harga_discount, &total)
				if err != nil {
					return nil, err
				}

				trxDataDetail, _ := transaction_detail.NewTransactionDetail(transaction_detail.DTOTransactionDetail{
					Id:              idTrxDetail,
					TransactionId:   transaction_id,
					ItemId:          item_id,
					JumlahPembelian: jumlah_pembelian,
					HargaPembelian:  harga_pembelian,
					HargaDiscount:   harga_discount,
					Total:           total,
				})

				trxDetailMany = append(trxDetailMany, trxDataDetail)
			}
			defer rowsTrxDetail.Close()

			// build struct transaction
			trxData, _ := transaction.NewTransactionWithDetail(transaction.DTOTransaction{
				Id:               idTrx,
				CustomerId:       customer_id,
				CodeTransaction:  code_transaction,
				Tanggalpembelian: &tanggal_pembelian,
				Total:            total,
			}, trxDetailMany)

			trxMany = append(trxMany, trxData)
		}
		defer rowsTrx.Close()

		// add to struct utama
		customer.AddTrx(trxMany)
		return customer, nil
	} else {
		return nil, nil
	}
}

func (repo *CustomerRepositoryMysqlInteractor) GetCustomerByNik(ctx context.Context, nik string) (*customer.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT * FROM customers where nik = ?"
	rows, errMysql := repo.dbConn.QueryContext(ctx, sqlQuery, nik)
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

		return customer, nil
	} else {
		return nil, nil
	}
}

func (repo *CustomerRepositoryMysqlInteractor) GetCustomerById(ctx context.Context, id int) (*customer.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	sqlQuery := "SELECT * FROM customers where id = ?"
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

		return customer, nil
	} else {
		return nil, nil
	}
}
