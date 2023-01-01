package model

type TransactionDetailModel struct {
	Id              int   `dbq:"id"`
	TransactionId   int   `dbq:"transaction_id"`
	ItemId          int   `dbq:"item_id"`
	JumlahPembelian int   `dbq:"jumlah_pembelian"`
	HargaPembelian  int64 `dbq:"harga_pembelian"`
	HargaDiscount   int64 `dbq:"harga_discount"`
	Total           int64 `dbq:"total"`
}

func GetTableTransactionDetail() string {
	return "transaction_detail"
}

func TableTransactionDetail() []string {
	return []string{
		"id",
		"transaction_id",
		"item_id",
		"jumlah_pembelian",
		"harga_pembelian",
		"harga_discount",
		"total",
	}
}
