package customer

import (
	"context"
	"database/sql"
	"game-store-final-project/project/domain/entity/customer"
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
