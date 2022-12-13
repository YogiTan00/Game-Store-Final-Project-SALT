package transaction

import (
	"context"
	"game-store-final-project/project/domain/entity/transaction"
	"game-store-final-project/project/domain/entity/transaction_detail"
	entity_voucher "game-store-final-project/project/domain/entity/voucher"
	"math/rand"
	"strconv"
	"time"
)

func (trx *TransactionUseCaseInteractor) UcGetAllTransaction(ctx context.Context) ([]*transaction.Transaction, error) {
	listTransaction, err := trx.repoTransaction.GetAllTransaction(ctx)
	if err != nil {
		return nil, err
	}

	return listTransaction, nil
}

func (trx *TransactionUseCaseInteractor) UcStoreTransaction(ctx context.Context, dataTransaction *transaction.DTOTransaction) error {
	trans, err := transaction.NewTransaction(*dataTransaction)
	if err != nil {
		return err
	}

	_, errInsert := trx.repoTransaction.StoreTransaction(ctx, trans)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (trx *TransactionUseCaseInteractor) UcGetAllTransactionByID(ctx context.Context, id string) ([]*transaction.Transaction, error) {
	listTransaction, err := trx.repoTransaction.GetAllTransactionByID(ctx, id)
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
		kupon 30% dengan prefix ULTI: ULTI-INV(code) 2022128(date) 7821387123456(13 rand digit)
		kupon 15% dengan prefix PREMI: PREMI-INV(code) 2022128(date) 7821387123456(13 rand digit)
		kupon 5% dengan prefix BASIC: BASIC-INV(code) 2022128(date) 7821387123456(13 rand digit)
	*/

	// get product
	var totalAksesorisAndNewGame int64 = 0
	var totalServiceConsole int64 = 0
	var totalSecondGame int64 = 0
	var totalSeluruh int64 = 0
	detailTransaction := make([]*transaction_detail.TransactionDetail, 0)

	for _, data_item := range items {
		dataItem, err := trx.repoItem.GetItemByID(trx.ctx, strconv.Itoa(data_item.ItemId))
		if err != nil {
			return nil, err
		}
		totalPerItem := int64(data_item.JumlahPembelian) * dataItem.GetHarga()

		if dataItem.GetKategori() == "Buy Accessories Console" || dataItem.GetKategori() == "Buy New Game" {
			totalAksesorisAndNewGame = totalAksesorisAndNewGame + totalPerItem
		}

		if dataItem.GetKategori() == "Service Console" {
			totalServiceConsole = totalServiceConsole + totalPerItem
		}

		if dataItem.GetKategori() == "Buy Second Game" {
			totalSecondGame = totalSecondGame + totalPerItem
		}

		totalSeluruh = totalSeluruh + totalPerItem

		// build dto details
		dataDetail, err := transaction_detail.NewTransactionDetailWithoutTrxId(transaction_detail.DTOTransactionDetail{
			ItemId:          dataItem.GetID(),
			JumlahPembelian: data_item.JumlahPembelian,
			HargaPembelian:  dataItem.GetHarga(),
			HargaDiscount:   0,
			Total:           totalPerItem,
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

	// update status voucher jika menggunakan

	return transaction, nil
}

func generateCodeVoucher(n int, name string, date time.Time) string {
	var randString = []rune("123456789")
	time := date.Format("20060102")
	b := make([]rune, n)
	for i := range b {
		b[i] = randString[rand.Intn(len(randString))]
	}
	return name + "-INV" + time + string(b)
}

func generateCodeTrx(n int, date time.Time) string {
	var randString = []rune("123456789")
	time := date.Format("20060102")
	b := make([]rune, n)
	for i := range b {
		b[i] = randString[rand.Intn(len(randString))]
	}
	return "INV-" + time + string(b)
}
