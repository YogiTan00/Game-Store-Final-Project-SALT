package customer

import (
	"context"
	"fmt"
	"game-store-final-project/project/domain/entity/customer"
	"game-store-final-project/project/internal/config/database/mysql"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mysqlConn         = mysql.InitMysqlDB()
	ctx               = context.Background()
	repoMysqlCustomer = NewCustomerRepositoryMysqlInteractor(mysqlConn)
)

func TestCustomerRepositoryMysqlInteractor_StoreCustomer(t *testing.T) {
	customer, _ := customer.NewCustomer(customer.DTOCustomer{
		Nik:          "3204223423442582",
		Nama:         "Taupik Pirdian",
		Alamat:       "Bandung",
		NoTlp:        "085846342122",
		JenisKelamin: "LK",
	})
	fmt.Println(customer)
	err := repoMysqlCustomer.StoreCustomer(ctx, customer)
	assert.Nil(t, err)
}

func TestCustomerRepositoryMysqlInteractor_IndexCustomerWithTransaction(t *testing.T) {
	customer, err := repoMysqlCustomer.GetCustomerByNik(ctx, "3241203322112233")
	fmt.Println(customer)
	assert.NotNil(t, customer)
	assert.Nil(t, err)
}

func TestCustomerRepositoryMysqlInteractor_GetCustomerByNik(t *testing.T) {
	customer, err := repoMysqlCustomer.GetCustomerByNik(ctx, "3241203322112233")
	fmt.Println(customer)
	assert.NotNil(t, customer)
	assert.Nil(t, err)
}

func TestCustomerRepositoryMysqlInteractor_GetCustomerById(t *testing.T) {
	customer, err := repoMysqlCustomer.GetCustomerById(ctx, 1)
	fmt.Println(customer)
	assert.NotNil(t, customer)
	assert.Nil(t, err)
}
