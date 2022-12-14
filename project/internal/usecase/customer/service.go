package customer

import (
	"game-store-final-project/project/domain/entity/customer"
)

func (cu *CustomerUseCaseInteractor) StoreCustomer(dataCustomer customer.DTOCustomer) (*customer.Customer, error) {
	// build data to entity
	customer, err := customer.NewCustomer(dataCustomer)
	if err != nil {
		return nil, err
	}

	// cek sudah ada data atau belum

	errInsert := cu.RepoCustomer.StoreCustomer(cu.ctx, customer)
	if errInsert != nil {
		return nil, errInsert
	}

	return customer, nil
}

func (cu *CustomerUseCaseInteractor) IndexCustomerTrx(nik string) (*customer.Customer, error) {
	// connect to repo
	customer, err := cu.RepoCustomer.IndexCustomerWithTransaction(cu.ctx, nik)
	if err != nil {
		return nil, err
	}

	return customer, nil
}
