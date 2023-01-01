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
	customer, err := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "3204223423442582",
		Nama:         "Taupik Pirdian",
		Alamat:       "Bandung",
		NoTlp:        "085846342122",
		JenisKelamin: "LK",
	})

	assert.Nil(t, err)
	assert.Equal(t, "3204223423442582", customer.GetNik())
}

/*
Negatif Case
*/
func TestValidationErrorNewCustomerNik(t *testing.T) {
	_, err := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "",
		Nama:         "Taupik Pirdian",
		Alamat:       "Bandung",
		NoTlp:        "085846342122",
		JenisKelamin: "LK",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "NIK NOT SET", err.Error())
}

func TestValidationErrorNewCustomerNikDigit(t *testing.T) {
	_, err := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "23",
		Nama:         "Taupik Pirdian",
		Alamat:       "Bandung",
		NoTlp:        "085846342122",
		JenisKelamin: "LK",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "NIK MUST 16 DIGIT", err.Error())
}

func TestValidationErrorNewCustomerNama(t *testing.T) {
	_, err := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "3204223423442582",
		Nama:         "",
		Alamat:       "Bandung",
		NoTlp:        "085846342122",
		JenisKelamin: "LK",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "NAMA NOT SET", err.Error())
}

func TestValidationErrorNewCustomerAlamat(t *testing.T) {
	_, err := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "3204223423442582",
		Nama:         "Taupik Pirdian",
		Alamat:       "",
		NoTlp:        "085846342122",
		JenisKelamin: "LK",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "ALAMAT NOT SET", err.Error())
}

func TestValidationErrorNewCustomerNoTelp(t *testing.T) {
	_, err := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "3204223423442582",
		Nama:         "Taupik Pirdian",
		Alamat:       "Bandung",
		NoTlp:        "",
		JenisKelamin: "LK",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "NO TLP NOT SET", err.Error())
}

func TestValidationErrorNewCustomerNoTelpAlphabet(t *testing.T) {
	_, err := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "3204223423442582",
		Nama:         "Taupik Pirdian",
		Alamat:       "Bandung",
		NoTlp:        "08584613241p",
		JenisKelamin: "LK",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "NO TLP CAN ONLY NUMBER", err.Error())
}

func TestValidationErrorNewCustomerJenisKelamin(t *testing.T) {
	_, err := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "3204223423442582",
		Nama:         "Taupik Pirdian",
		Alamat:       "Bandung",
		NoTlp:        "085846342122",
		JenisKelamin: "",
	})

	assert.NotNil(t, err)
	assert.Equal(t, "JENIS KELAMIN NOT SET", err.Error())
}
