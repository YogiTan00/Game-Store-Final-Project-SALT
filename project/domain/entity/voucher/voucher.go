package voucher

import (
	"errors"
	"time"
)

type Voucher struct {
	id            int
	customerId    int
	transactionId int
	codeVoucher   string
	namaVoucher   string
	startDate     time.Time
	endDate       time.Time
	useDate       time.Time
	status        int
	nilaiDisc     int
	typeDisc      int
}

type DTOVoucher struct {
	Id            int
	CustomerId    int
	TransactionId int
	CodeVoucher   string
	NamaVoucher   string
	StartDate     time.Time
	EndDate       time.Time
	UseDate       time.Time
	Status        int
	NilaiDisc     int
	TypeDisc      int
}

func NewVoucher(t []*DTOVoucher) ([]*Voucher, error) {
	vouchers := make([]*Voucher, 0)
	for _, voucher := range t {
		addVoucher := &Voucher{
			customerId:  voucher.CustomerId,
			codeVoucher: voucher.CodeVoucher,
			namaVoucher: voucher.NamaVoucher,
			startDate:   voucher.StartDate,
			endDate:     voucher.EndDate,
			useDate:     voucher.UseDate,
			status:      voucher.Status,
			nilaiDisc:   voucher.NilaiDisc,
			typeDisc:    voucher.TypeDisc,
		}
		vouchers = append(vouchers, addVoucher)
	}

	return vouchers, nil
}

func NewVoucherSingle(v DTOVoucher) (*Voucher, error) {
	if v.CustomerId == 0 {
		return nil, errors.New("CUSTOMER ID NOT SET")
	}

	if v.CodeVoucher == "" {
		return nil, errors.New("CODE NOT SET")
	}

	if v.NamaVoucher == "" {
		return nil, errors.New("NAME NOT SET")
	}

	if v.StartDate.IsZero() {
		return nil, errors.New("START Date NOT SET")
	}

	if v.EndDate.IsZero() {
		return nil, errors.New("END Date NOT SET")
	}

	if v.Status == 0 {
		return nil, errors.New("STATUS NOT SET")
	}

	if v.NilaiDisc == 0 {
		return nil, errors.New("NILAI NOT SET")
	}

	if v.TypeDisc == 0 {
		return nil, errors.New("TYPE NOT SET")
	}

	voucher := &Voucher{
		id:            v.Id,
		customerId:    v.CustomerId,
		transactionId: v.TransactionId,
		codeVoucher:   v.CodeVoucher,
		namaVoucher:   v.NamaVoucher,
		startDate:     v.StartDate,
		endDate:       v.EndDate,
		useDate:       v.UseDate,
		status:        v.Status,
		nilaiDisc:     v.NilaiDisc,
		typeDisc:      v.TypeDisc,
	}

	return voucher, nil
}

func NewVoucherCustomer(v DTOVoucher) (*Voucher, error) {
	if v.CustomerId == 0 {
		return nil, errors.New("CUSTOMER ID NOT SET")
	}

	if v.CodeVoucher == "" {
		return nil, errors.New("CODE NOT SET")
	}

	if v.NamaVoucher == "" {
		return nil, errors.New("NAME NOT SET")
	}

	if v.StartDate.IsZero() {
		return nil, errors.New("START Date NOT SET")
	}

	if v.EndDate.IsZero() {
		return nil, errors.New("END Date NOT SET")
	}

	if v.NilaiDisc == 0 {
		return nil, errors.New("NILAI NOT SET")
	}

	if v.TypeDisc == 0 {
		return nil, errors.New("TYPE NOT SET")
	}

	voucher := &Voucher{
		id:            v.Id,
		customerId:    v.CustomerId,
		transactionId: v.TransactionId,
		codeVoucher:   v.CodeVoucher,
		namaVoucher:   v.NamaVoucher,
		startDate:     v.StartDate,
		endDate:       v.EndDate,
		useDate:       v.UseDate,
		status:        v.Status,
		nilaiDisc:     v.NilaiDisc,
		typeDisc:      v.TypeDisc,
	}

	return voucher, nil
}

func (vo *Voucher) GetId() int {
	return vo.id
}

func (vo *Voucher) GetCustomerId() int {
	return vo.customerId
}

func (vo *Voucher) GetCodeVoucher() string {
	return vo.codeVoucher
}

func (vo *Voucher) GetNamaVoucher() string {
	return vo.namaVoucher
}

func (vo *Voucher) GetStartDate() time.Time {
	return vo.startDate
}

func (vo *Voucher) GetEndDate() time.Time {
	return vo.endDate
}

func (vo *Voucher) GetUseDate() time.Time {
	return vo.useDate
}

func (vo *Voucher) GetStatus() int {
	return vo.status
}

func (vo *Voucher) GetNilaiDisc() int {
	return vo.nilaiDisc
}

func (vo *Voucher) GetTypeDisc() int {
	return vo.typeDisc
}
