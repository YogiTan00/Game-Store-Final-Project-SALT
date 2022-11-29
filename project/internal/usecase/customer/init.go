package customer

import (
	"context"
	"game-store-final-project/project/domain/repository"
)

type CustomerUseCaseInteractor struct {
	ctx          context.Context
	RepoCustomer repository.CustomerRepository
}

// ini ngeimplement domain usecase
func NewCustomerUseCaseInteractor(ctx context.Context, repoImplement repository.CustomerRepository) *CustomerUseCaseInteractor {
	return &CustomerUseCaseInteractor{ctx: ctx, RepoCustomer: repoImplement}
}
