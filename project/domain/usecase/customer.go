package usecase

import (
	"game-store-final-project/project/domain/entity/customer"
)

/*
usecase adalah aturan logic dari team bisnis
*/

type CustomerUseCase interface {
	StoreCustomer(nik string, nama string, alamat string, no_tlp string, jenis_kelamin string) (*customer.Customer, error)
}
