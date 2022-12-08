package http_request

type RequestTransaction struct {
	CustomerId        int                  `json:"customer_id"`
	TanggalPembelian  string               `json:"tanggal_pembelian"`
	Voucher           []string             `json:"voucher"`
	DetailTransaction *[]DetailTransaction `json:"detail_transaction"`
}

type DetailTransaction struct {
	ItemId          int `json:"item_id"`
	JumlahPembelian int `json:"jumlah_pembelian"`
}
