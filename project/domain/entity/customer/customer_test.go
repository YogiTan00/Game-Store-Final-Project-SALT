package customer_test

import (
	"game-store-final-project/project/domain/entity/customer"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Positif Case
*/
func TestNewCustomer(t *testing.T) {
	customer, err := customer.NewCustomer(
		"3204223423442582",
		"Taupik Pirdian",
		"Bandung",
		"085846342122",
		"LK",
	)

	assert.Nil(t, err)
	assert.Equal(t, "3204223423442582", customer.GetNik())
}

/*
Negatif Case
*/
func TestValidationErrorNewCustomerNik(t *testing.T) {
	_, err := customer.NewCustomer(
		"",
		"Taupik Pirdian",
		"Bandung",
		"085846342122",
		"LK",
	)

	assert.NotNil(t, err)
	assert.Equal(t, "NIK NOT SET", err.Error())
}

func TestValidationErrorNewCustomerNama(t *testing.T) {
	_, err := customer.NewCustomer(
		"3204223423442582",
		"",
		"Bandung",
		"085846342122",
		"LK",
	)

	assert.NotNil(t, err)
	assert.Equal(t, "NAMA NOT SET", err.Error())
}

func TestValidationErrorNewCustomerAlamat(t *testing.T) {
	_, err := customer.NewCustomer(
		"3204223423442582",
		"Taupik Pirdian",
		"",
		"085846342122",
		"LK",
	)

	assert.NotNil(t, err)
	assert.Equal(t, "ALAMAT NOT SET", err.Error())
}

func TestValidationErrorNewCustomerNoTelp(t *testing.T) {
	_, err := customer.NewCustomer(
		"3204223423442582",
		"Taupik Pirdian",
		"Bandung",
		"",
		"LK",
	)

	assert.NotNil(t, err)
	assert.Equal(t, "NO TLP NOT SET", err.Error())
}

func TestValidationErrorNewCustomerJenisKelamin(t *testing.T) {
	_, err := customer.NewCustomer(
		"3204223423442582",
		"Taupik Pirdian",
		"Bandung",
		"085846342122",
		"",
	)

	assert.NotNil(t, err)
	assert.Equal(t, "JENIS KELAMIN NOT SET", err.Error())
}
