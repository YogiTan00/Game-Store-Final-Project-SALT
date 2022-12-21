package transaction

import (
	"context"
	"errors"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/domain/entity/transaction_detail"
	entity_voucher "game-store-final-project/project/domain/entity/voucher"
	"math/rand"
	"strconv"
	"time"
)

func (trx *TransactionUseCaseInteractor) GetTransactionByID(ctx context.Context, id string) (*transaction.Transaction, error) {
	listTransaction, err := trx.repoTransaction.GetTransactionByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return listTransaction, nil
}

func (trx *TransactionUseCaseInteractor) GetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error) {
	listTransaction, err := trx.repoTransaction.GetAllTransaction(ctx)
	if err != nil {
		return nil, err
	}

	return listTransaction, nil
}

func (trx *TransactionUseCaseInteractor) GetAllTransactionByCustomerID(ctx context.Context, id string) ([]*transaction.Transaction, error) {
	listTransaction, err := trx.repoTransaction.GetAllTransactionByCustomerID(ctx, id)
	if err != nil {
		return nil, err
	}

	return listTransaction, nil
}

func (trx *TransactionUseCaseInteractor) StoreTransaction(customer_id int, tanggal_pembelian string, voucher []string, items []*transaction.DTOItemPembelian) (*transaction.Transaction, error) {
	/*
		Rule 1:
		cek produknya ready atau tidak, jika iya return datanya

		Rule 2:
		cek dapet diskon atau tidak,
		jika iya generate ke table discount dan return datanya

		Rule Discount:
		Pelanggan yang mempunyai revenuenya more than 25.000.000 give 30% discount (Accessories & New Game)
		Pelanggan yang mempunyai revenuenya melebihi 13.000.000 get disc 15% (Service Console)
		Pelanggan yang purchase more than 6.000.000 get disc 5% (Second Game)

		Rule CodeVocuher:
		kupon 30% dengan prefix ULTI: ULTI-2022Y12M08D12H02M12S(date) 78SD138SS1234(13 rand digit)
		kupon 15% dengan prefix PREMI: PREMI-2022Y12M08D12H02M12S(date) 78SD138SS1234(13 rand digit)
		kupon 5% dengan prefix BASIC: BASIC-2022Y12M08D12H02M12S(date) 78SD138SS1234(13 rand digit)
	*/

	// check user terdaftar atau tidak
	customer, errCustomer := trx.repoCustomer.GetCustomerById(trx.ctx, customer_id)
	if errCustomer != nil {
		return nil, errCustomer
	}
	if customer == nil {
		return nil, errors.New("CUSTOMER NOT FOUND")
	}

	// check voucher jika menggunakan
	var discountAksesorisAndNewGame float64 = 0
	var discountServiceConsole float64 = 0
	var discountSecondGame float64 = 0

	if len(voucher) > 0 {
		for _, v := range voucher {
			dataVoucher, errVocuher := trx.repoVoucher.GetVoucherByCode(trx.ctx, v)
			if errVocuher != nil {
				return nil, errVocuher
			}

			if dataVoucher == nil {
				return nil, errors.New("VOUCHER NOT FOUND OR NON ACTIVE")
			} else {
				// update status voucher jika menggunakan
				errUpdate := trx.repoVoucher.UpdateVoucherById(trx.ctx, dataVoucher.GetId())
				if errUpdate != nil {
					return nil, errUpdate
				}
			}

			if dataVoucher.GetNamaVoucher() == "ULTI" {
				discountAksesorisAndNewGame = discountAksesorisAndNewGame + float64(dataVoucher.GetNilaiDisc())
			}

			if dataVoucher.GetNamaVoucher() == "PREMI" {
				discountServiceConsole = discountServiceConsole + float64(dataVoucher.GetNilaiDisc())
			}

			if dataVoucher.GetNamaVoucher() == "BASIC" {
				discountSecondGame = discountSecondGame + float64(dataVoucher.GetNilaiDisc())
			}
		}
	}

	// get product
	var totalAksesorisAndNewGame int64 = 0
	var totalServiceConsole int64 = 0
	var totalSecondGame int64 = 0
	var totalSeluruh int64 = 0

	detailTransaction := make([]*transaction_detail.TransactionDetail, 0)
	for _, data_item := range items {
		var discountPrice float64 = 0

		dataItem, err := trx.repoItem.GetItemByID(trx.ctx, strconv.Itoa(data_item.ItemId))
		if err != nil {
			return nil, err
		}
		totalPerItem := int64(data_item.JumlahPembelian) * dataItem.GetHarga()

		if dataItem.GetKategori() == "Buy Accessories Console" || dataItem.GetKategori() == "Buy New Game" {
			totalAksesorisAndNewGame = totalAksesorisAndNewGame + totalPerItem
			if discountAksesorisAndNewGame > 0 {
				discountPrice = (discountAksesorisAndNewGame / 100) * float64(totalAksesorisAndNewGame)
			}
		}

		if dataItem.GetKategori() == "Service Console" {
			totalServiceConsole = totalServiceConsole + totalPerItem
			if discountServiceConsole > 0 {
				discountPrice = (discountServiceConsole / 100) * float64(totalServiceConsole)
			}
		}

		if dataItem.GetKategori() == "Buy Second Game" {
			totalSecondGame = totalSecondGame + totalPerItem
			if discountSecondGame > 0 {
				discountPrice = (discountSecondGame / 100) * float64(totalSecondGame)
			}
		}

		valueDiscount := int64(discountPrice)
		totalSeluruh = (totalSeluruh + totalPerItem) - valueDiscount

		// build dto details
		dataDetail, err := transaction_detail.NewTransactionDetailWithoutTrxId(transaction_detail.DTOTransactionDetail{
			ItemId:          dataItem.GetID(),
			JumlahPembelian: data_item.JumlahPembelian,
			HargaPembelian:  dataItem.GetHarga(),
			HargaDiscount:   valueDiscount,
			Total:           totalPerItem - valueDiscount,
		})

		detailTransaction = append(detailTransaction, dataDetail)
	}

	dateNow := time.Now()
	getVouchers := make([]*entity_voucher.DTOVoucher, 0)
	if totalAksesorisAndNewGame > 25000000 {
		codeVoucher := "ULTI"
		fullCodeVoucher := generateCodeVoucher(13, codeVoucher, dateNow)

		accessoriesVoucher := &entity_voucher.DTOVoucher{
			CustomerId:  customer_id,
			CodeVoucher: fullCodeVoucher,
			NamaVoucher: "ULTI",
			StartDate:   dateNow,
			EndDate:     dateNow.AddDate(1, 0, 0),
			Status:      1,
			NilaiDisc:   30,
			TypeDisc:    1,
		}
		getVouchers = append(getVouchers, accessoriesVoucher)
	}

	if totalServiceConsole > 13000000 {
		codeVoucher := "PREMI"
		fullCodeVoucher := generateCodeVoucher(13, codeVoucher, dateNow)

		serviceVoucher := &entity_voucher.DTOVoucher{
			CustomerId:  customer_id,
			CodeVoucher: fullCodeVoucher,
			NamaVoucher: "PREMI",
			StartDate:   dateNow,
			EndDate:     dateNow.AddDate(1, 0, 0),
			Status:      1,
			NilaiDisc:   15,
			TypeDisc:    2,
		}
		getVouchers = append(getVouchers, serviceVoucher)
	}

	if totalSecondGame > 6000000 {
		codeVoucher := "BASIC"
		fullCodeVoucher := generateCodeVoucher(13, codeVoucher, dateNow)

		secondVoucher := &entity_voucher.DTOVoucher{
			CustomerId:  customer_id,
			CodeVoucher: fullCodeVoucher,
			NamaVoucher: "BASIC",
			StartDate:   dateNow,
			EndDate:     dateNow.AddDate(1, 0, 0),
			Status:      1,
			NilaiDisc:   5,
			TypeDisc:    3,
		}
		getVouchers = append(getVouchers, secondVoucher)
	}

	// build data to entity voucher
	dataVoucher, errVoucher := entity_voucher.NewVoucher(getVouchers)
	if errVoucher != nil { // error build voucher
		return nil, errVoucher
	}

	// store voucher
	// loop and save
	for _, list_vo := range dataVoucher {
		errInsertVo := trx.repoVoucher.StoreVoucher(trx.ctx, list_vo)
		if errInsertVo != nil {
			return nil, errInsertVo
		}
	}

	tglBeli, _ := time.Parse("2006-01-02", tanggal_pembelian)
	// build data to entity transaction
	transaction, err := transaction.NewTransaction(transaction.DTOTransaction{
		CustomerId:       customer_id,
		CodeTransaction:  generateCodeTrx(13, dateNow),
		Tanggalpembelian: &tglBeli,
		Total:            totalSeluruh,
	})

	if err != nil { // error build trx
		return nil, err
	}

	// store trx
	resStore, errInsert := trx.repoTransaction.StoreTransaction(trx.ctx, transaction)
	if errInsert != nil {
		return nil, errInsert
	}

	// store trx detail
	for _, listDetail := range detailTransaction {
		errInsertDetail := trx.repoTransactionDetail.StoreTransactionDetail(trx.ctx, resStore, listDetail)
		if errInsertDetail != nil {
			return nil, errInsertDetail
		}
	}

	return transaction, nil
}

func generateCodeVoucher(n int, name string, date time.Time) string {
	var randString = []rune("123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	time := date.Format("02D01M2006Y15H04M05S")
	b := make([]rune, n)
	for i := range b {
		b[i] = randString[rand.Intn(len(randString))]
	}
	return name + "-" + time + string(b)
}

func generateCodeTrx(n int, date time.Time) string {
	var randString = []rune("123456789")
	time := date.Format("02D01M2006Y15H04M05S")
	b := make([]rune, n)
	for i := range b {
		b[i] = randString[rand.Intn(len(randString))]
	}
	return "INV-" + time
}
