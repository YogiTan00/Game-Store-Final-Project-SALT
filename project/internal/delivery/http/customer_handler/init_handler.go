package customer_handler

import (
	"context"
	_repository "game-store-final-project/project/domain/repository"
)

type CustomerHandler struct {
	ctx          context.Context
	repoCustomer _repository.CustomerRepository
}

func NewCustomerHandler(ctx context.Context, repoCustomer _repository.CustomerRepository) *CustomerHandler {
	return &CustomerHandler{
		ctx:          ctx,
		repoCustomer: repoCustomer,
	}
}
