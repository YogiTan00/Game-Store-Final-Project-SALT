package usecase

import (
	"game-store-final-project/project/domain/entity/customer"
)

/*
usecase adalah aturan logic dari team bisnis
*/

type CustomerUseCase interface {
	StoreCustomer(dc customer.DTOCustomer) (*customer.Customer, error)
	IndexCustomerWithTransaction(nik string) (*customer.Customer, error)
}
