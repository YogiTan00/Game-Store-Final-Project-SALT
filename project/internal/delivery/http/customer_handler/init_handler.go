package customer_handler

import (
	"game-store-final-project/project/domain/usecase"
)

/*
init handler ini untuk mengarahkan ke usecase
*/
type CustomerHandlerInteractor struct {
	CustomerUseCase usecase.CustomerUseCase
}

func NewCustomerHandler(usecaseImplement usecase.CustomerUseCase) *CustomerHandlerInteractor {
	return &CustomerHandlerInteractor{CustomerUseCase: usecaseImplement}
}
