package customer

import (
	"game-store-final-project/project/domain/entity/customer"
)

func (r *RepoCustomer) StoreCustomer(dc customer.DTOCustomer) (*customer.Customer, error) {
	args := r.Called(dc)
	return args.Get(0).(*customer.Customer), args.Error(0)
}

func (r *RepoCustomer) IndexCustomerWithTransaction(nik string) (*customer.Customer, error) {
	args := r.Called(nik)
	return args.Get(0).(*customer.Customer), args.Error(1)
}
