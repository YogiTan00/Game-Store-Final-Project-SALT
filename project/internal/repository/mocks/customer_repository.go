package mocks

import (
	"context"
	"game-store-final-project/project/domain/entity/customer"

	"github.com/stretchr/testify/mock"
)

type CustomerRepositoryMock struct {
	mock.Mock
}

func (cum *CustomerRepositoryMock) GetCustomerByNik(ctx context.Context, nik string) ([]*customer.Customer, error) {
	args := cum.Called(ctx, nik)
	return args.Get(0).([]*customer.Customer), args.Error(1)
}

func (cum *CustomerRepositoryMock) GetCustomerById(ctx context.Context, id int) ([]*customer.Customer, error) {
	args := cum.Called(ctx, id)
	return args.Get(0).([]*customer.Customer), args.Error(1)
}
