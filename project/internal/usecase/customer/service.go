package customer

import (
	"game-store-final-project/project/domain/entity"
)

func (cu *CustomerUseCaseInteractor) StoreCustomer(nik string, nama string, alamat string, no_tlp string, jenis_kelamin string) (*entity.Customer, error) {
	// build data to entity
	customer, err := entity.NewCustomer(nik, nama, alamat, no_tlp, jenis_kelamin)
	if err != nil {
		return nil, err
	}

	errInsert := cu.RepoCustomer.StoreCustomer(cu.ctx, customer)
	if errInsert != nil {
		return nil, errInsert
	}

	return customer, nil
}
