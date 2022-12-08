package voucher

import "time"

type Voucher struct {
	customerId  int
	codeVoucher string
	namaVoucher string
	startDate   time.Time
	endDate     time.Time
	useDate     time.Time
	status      int
	nilaiDisc   int
	typeDisc    int
}

type DTOVoucher struct {
	CustomerId  int
	CodeVoucher string
	NamaVoucher string
	StartDate   time.Time
	EndDate     time.Time
	UseDate     time.Time
	Status      int
	NilaiDisc   int
	TypeDisc    int
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
