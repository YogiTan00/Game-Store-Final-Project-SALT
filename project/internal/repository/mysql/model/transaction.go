package model

import "time"

type TransactionModel struct {
	Id               int        `dbq:"id"`
	CustomerId       int        `dbq:"customer_id"`
	CodeTransaction  string     `dbq:"code_transaction"`
	TanggalPembelian *time.Time `dbq:"tanggal_pembelian"`
	Total            int64      `dbq:"total"`
}

func GetTableTransaction() string {
	return "transactions"
}

func TableTransaction() []string {
	return []string{
		"id",
		"customer_id",
		"code_transaction",
		"tanggal_pembelian",
		"total",
	}
}
