package customer

import (
	"context"
	"game-store-final-project/project/internal/repository/mocks"
)

var (
	repoCustomerMock = new(mocks.CustomerRepositoryMock)
	ctx              = context.Background()
)

// func TestIndexCustomerWithTransaction(t *testing.T) {
// 	repoCustomerMock.On("IndexCustomerWithTransaction", mock.Anything, mock.Anything).Return(test_data.DataCustomerWithTransaction(), nil)

// 	useCase := NewCustomerUseCaseInteractor(ctx, repoCustomerMock)
// 	listArticle, errGetArticle := useCase.IndexCustomerWithTransaction(mock.Anything)
// 	assert.Nil(t, errGetArticle)
// 	assert.NotNil(t, listArticle)
// }
