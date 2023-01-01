package http_request

type RequestCustomer struct {
	Nik          string `json:"nik"`
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	NoTlp        string `json:"no_tlp"`
	JenisKelamin string `json:"jenis_kelamin"`
}
