package model

type TransactionDetailModel struct {
	Id              int    `dbq:"id"`
	CodeTrans       string `dbq:"code_transaction"`
	TransactionId   int    `dbq:"transaction_id"`
	ItemId          string `dbq:"item_id"`
	JumlahPembelian int    `dbq:"jumlah_pembelian"`
	HargaPembelian  int64  `dbq:"harga_pembelian"`
	Total           int64  `dbq:"total"`
}

func GetTableTransactionDetail() string {
	return "transaction_detail"
}

func TableTransactionDetail() []string {
	return []string{
		"id",
		"code_transaction",
		"transaction_id",
		"item_id",
		"jumlah_pembelian",
		"harga_pembelian",
		"total",
	}
}
