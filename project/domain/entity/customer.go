package entity

import "errors"

type Customer struct {
	nik           string
	nama          string
	alamat        string
	no_tlp        string
	jenis_kelamin string
}

func NewCustomer(nik string, nama string, alamat string, no_tlp string, jenis_kelamin string) (*Customer, error) {
	if nik == "" {
		return nil, errors.New("NIK NOT SET")
	}

	if nama == "" {
		return nil, errors.New("NAMA NOT SET")
	}

	if alamat == "" {
		return nil, errors.New("ALAMAT NOT SET")
	}

	if no_tlp == "" {
		return nil, errors.New("NO TLP NOT SET")
	}

	if jenis_kelamin == "" {
		return nil, errors.New("JENIS KELAMIN NOT SET")
	}

	return &Customer{
		nik:           nik,
		nama:          nama,
		alamat:        alamat,
		no_tlp:        no_tlp,
		jenis_kelamin: jenis_kelamin,
	}, nil
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
	return cu.no_tlp
}

func (cu *Customer) GetJenisKelamin() string {
	return cu.jenis_kelamin
}
