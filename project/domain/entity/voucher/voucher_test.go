package voucher_test

import (
	"game-store-final-project/project/domain/entity/voucher"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*
Positif Case
*/
func TestNewVoucher(t *testing.T) {
	_, err := voucher.NewVoucher([]*voucher.DTOVoucher{
		{
			CustomerId:    1,
			TransactionId: 1,
			CodeVoucher:   "PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M6",
			NamaVoucher:   "PREMI",
			StartDate:     time.Now(),
			EndDate:       time.Now(),
			UseDate:       time.Now(),
			Status:        1,
			NilaiDisc:     15,
			TypeDisc:      2,
		},
		{
			CustomerId:    1,
			TransactionId: 1,
			CodeVoucher:   "PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M6",
			NamaVoucher:   "PREMI",
			StartDate:     time.Now(),
			EndDate:       time.Now(),
			UseDate:       time.Now(),
			Status:        1,
			NilaiDisc:     15,
			TypeDisc:      2,
		},
	})

	assert.Nil(t, err)
}

func TestNewVoucherSingle(t *testing.T) {
	voucher, err := voucher.NewVoucherSingle(voucher.DTOVoucher{
		CustomerId:    1,
		TransactionId: 1,
		CodeVoucher:   "PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M6",
		NamaVoucher:   "PREMI",
		StartDate:     time.Now(),
		EndDate:       time.Now(),
		UseDate:       time.Now(),
		Status:        1,
		NilaiDisc:     15,
		TypeDisc:      2,
	})

	assert.Nil(t, err)
	assert.Equal(t, "PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M6", voucher.GetCodeVoucher())
}

/*
Negatif Case
*/
func TestValidationErrorNewVoucherCustomerId(t *testing.T) {
	_, err := voucher.NewVoucherSingle(voucher.DTOVoucher{
		CustomerId:    0,
		TransactionId: 1,
		CodeVoucher:   "PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M6",
		NamaVoucher:   "PREMI",
		StartDate:     time.Now(),
		EndDate:       time.Now(),
		UseDate:       time.Now(),
		Status:        1,
		NilaiDisc:     15,
		TypeDisc:      2,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "CUSTOMER ID NOT SET", err.Error())
}

func TestValidationErrorNewVoucherCode(t *testing.T) {
	_, err := voucher.NewVoucherSingle(voucher.DTOVoucher{
		CustomerId:    1,
		TransactionId: 1,
		CodeVoucher:   "",
		NamaVoucher:   "PREMI",
		StartDate:     time.Now(),
		EndDate:       time.Now(),
		UseDate:       time.Now(),
		Status:        1,
		NilaiDisc:     15,
		TypeDisc:      2,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "CODE NOT SET", err.Error())
}

func TestValidationErrorNewName(t *testing.T) {
	_, err := voucher.NewVoucherSingle(voucher.DTOVoucher{
		CustomerId:    1,
		TransactionId: 1,
		CodeVoucher:   "PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M6",
		NamaVoucher:   "",
		StartDate:     time.Now(),
		EndDate:       time.Now(),
		UseDate:       time.Now(),
		Status:        1,
		NilaiDisc:     15,
		TypeDisc:      2,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "NAME NOT SET", err.Error())
}

func TestValidationErrorNewStatus(t *testing.T) {
	_, err := voucher.NewVoucherSingle(voucher.DTOVoucher{
		CustomerId:    1,
		TransactionId: 1,
		CodeVoucher:   "PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M6",
		NamaVoucher:   "PREMI",
		StartDate:     time.Now(),
		EndDate:       time.Now(),
		UseDate:       time.Now(),
		Status:        0,
		NilaiDisc:     15,
		TypeDisc:      2,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "STATUS NOT SET", err.Error())
}

func TestValidationErrorNewNilai(t *testing.T) {
	_, err := voucher.NewVoucherSingle(voucher.DTOVoucher{
		CustomerId:    1,
		TransactionId: 1,
		CodeVoucher:   "PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M6",
		NamaVoucher:   "PREMI",
		StartDate:     time.Now(),
		EndDate:       time.Now(),
		UseDate:       time.Now(),
		Status:        1,
		NilaiDisc:     0,
		TypeDisc:      2,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "NILAI NOT SET", err.Error())
}

func TestValidationErrorNewType(t *testing.T) {
	_, err := voucher.NewVoucherSingle(voucher.DTOVoucher{
		CustomerId:    1,
		TransactionId: 1,
		CodeVoucher:   "PREMI-20D12M2022Y14H16M17SNEQVM4SCP78M6",
		NamaVoucher:   "PREMI",
		StartDate:     time.Now(),
		EndDate:       time.Now(),
		UseDate:       time.Now(),
		Status:        1,
		NilaiDisc:     15,
		TypeDisc:      0,
	})

	assert.NotNil(t, err)
	assert.Equal(t, "TYPE NOT SET", err.Error())
}
