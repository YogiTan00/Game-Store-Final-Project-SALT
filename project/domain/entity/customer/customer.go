package customer

import (
	"errors"
	"game-store-final-project/project/domain/entity/transaction"
)

type Customer struct {
	id           int
	nik          string
	nama         string
	alamat       string
	noTlp        string
	jenisKelamin string
	transaction  []*transaction.Transaction
}

type DTOCustomer struct {
	Id           int
	Nik          string
	Nama         string
	Alamat       string
	NoTlp        string
	JenisKelamin string
	Transaction  []*transaction.DTOTransaction
}

/*
Ganti pake dto
*/
func NewCustomer(dataCustomer DTOCustomer) (*Customer, error) {
	if dataCustomer.Nik == "" {
		return nil, errors.New("NIK NOT SET")
	}

	if dataCustomer.Nama == "" {
		return nil, errors.New("NAMA NOT SET")
	}

	if dataCustomer.Alamat == "" {
		return nil, errors.New("ALAMAT NOT SET")
	}

	if dataCustomer.NoTlp == "" {
		return nil, errors.New("NO TLP NOT SET")
	}

	if dataCustomer.JenisKelamin == "" {
		return nil, errors.New("JENIS KELAMIN NOT SET")
	}

	return &Customer{
		id:           dataCustomer.Id,
		nik:          dataCustomer.Nik,
		nama:         dataCustomer.Nama,
		alamat:       dataCustomer.Alamat,
		noTlp:        dataCustomer.NoTlp,
		jenisKelamin: dataCustomer.JenisKelamin,
	}, nil
}

func (cu *Customer) AddTrx(dataTransaction []*transaction.Transaction) *Customer {
	cu.transaction = dataTransaction
	return cu
}

func (cu *Customer) GetId() int {
	return cu.id
}

func (cu *Customer) GetNik() string {
	return cu.nik
}

func (cu *Customer) GetNama() string {
	return cu.nama
}

func (cu *Customer) GetAlamat() string {
	return cu.alamat
}

func (cu *Customer) GetNoTlp() string {
	return cu.noTlp
}

func (cu *Customer) GetJenisKelamin() string {
	return cu.jenisKelamin
}

func (cu *Customer) GetDetailTrx() []*transaction.Transaction {
	return *&cu.transaction
}
