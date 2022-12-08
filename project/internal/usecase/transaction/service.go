package transaction

import (
	"fmt"
	"game-store-final-project/project/domain/entity/transaction"
	entity_voucher "game-store-final-project/project/domain/entity/voucher"
	"strconv"
	"time"
)

func (cu *TransactionUseCaseInteractor) StoreTransaction(customer_id int, tanggal_pembelian string, voucher []string, items []*transaction.DTOItemPembelian) (*transaction.Transaction, error) {
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
		kupon 30% dengan prefix ULTI: ULTI-RND7821387123456(16 digit)
		kupon 15% dengan prefix PREMI: PREMI-RND1209318092312(16 digit)
		kupon 5% dengan prefix BASIC: BASIC-RND1923808132345(16 digit)
	*/

	// get product
	var totalSeluruh int64 = 0
	for _, data_item := range items {
		dataItem, err := cu.repoItem.GetItemByID(cu.ctx, strconv.Itoa(data_item.ItemId))
		if err != nil {
			return nil, err
		}

		totalPerItem := int64(data_item.JumlahPembelian) * dataItem.GetHarga()
		totalSeluruh = totalSeluruh + totalPerItem
	}

	dateNow := time.Now()
	getVouchers := make([]*entity_voucher.DTOVoucher, 0)
	if totalSeluruh > 25000000 {
		codeVoucher := "ULTI"
		accessoriesVoucher := &entity_voucher.DTOVoucher{
			CustomerId:  customer_id,
			CodeVoucher: codeVoucher,
			NamaVoucher: "ULTI",
			StartDate:   dateNow,
			EndDate:     dateNow.AddDate(1, 0, 0),
			Status:      1,
			NilaiDisc:   30,
			TypeDisc:    1,
		}
		getVouchers = append(getVouchers, accessoriesVoucher)
	}

	if totalSeluruh > 13000000 {
		codeVoucher := "PREMI"
		serviceVoucher := &entity_voucher.DTOVoucher{
			CustomerId:  customer_id,
			CodeVoucher: codeVoucher,
			NamaVoucher: "PREMI",
			StartDate:   dateNow,
			EndDate:     dateNow.AddDate(1, 0, 0),
			Status:      1,
			NilaiDisc:   15,
			TypeDisc:    1,
		}
		getVouchers = append(getVouchers, serviceVoucher)
	}

	if totalSeluruh > 6000000 {
		codeVoucher := "BASIC"
		secondVoucher := &entity_voucher.DTOVoucher{
			CustomerId:  customer_id,
			CodeVoucher: codeVoucher,
			NamaVoucher: "BASIC",
			StartDate:   dateNow,
			EndDate:     dateNow.AddDate(0, 0, 7),
			Status:      1,
			NilaiDisc:   5,
			TypeDisc:    1,
		}
		getVouchers = append(getVouchers, secondVoucher)
	}

	// build data to entity voucher
	vocuher, errVoucher := entity_voucher.NewVoucher(getVouchers)
	if errVoucher != nil { // error build voucher
		return nil, errVoucher
	}

	fmt.Println(vocuher)

	// build data to entity transaction
	transaction, err := transaction.NewTransaction(transaction.DTOTransaction{
		VoucherCustomerId: 1,
		CustomerId:        customer_id,
		CodeTransaction:   "",
		Tanggalpembelian:  tanggal_pembelian,
		Total:             0,
		HargaDiscount:     0,
		TotalHarga:        0,
		ItemPembelian:     items,
	})

	if err != nil { // error build trx
		return nil, err
	}

	// store trx
	errInsert := cu.repoTransaction.StoreTransaction(cu.ctx, transaction)
	if errInsert != nil {
		return nil, errInsert
	}

	// store voucher

	return transaction, nil
}
