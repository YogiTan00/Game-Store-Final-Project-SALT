package entity_test

import (
	"game-store-final-project/project/domain/entity"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

/*
Positif Case
*/
func TestNewTransaction(t *testing.T) {
	transaction, err := entity.NewTransaction(entity.DTOTransaction{
		CustomerId:      25123123,
		CodeTransaction: time.Now(),
	})

	assert.Nil(t, err)
	assert.Equal(t, 25123123, transaction.GetCustomerID())
}

/*
Negative Case
*/

func TestValidateTransactionID(t *testing.T) {
	_, err := entity.NewTransaction(entity.DTOTransaction{
		CustomerId:      0,
		CodeTransaction: time.Now(),
	})

	assert.NotNil(t, err)
	assert.Equal(t, "ID COSTOMER NOT SET", err.Error())
}
