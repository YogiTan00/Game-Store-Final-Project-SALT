package model

type ItemModel struct {
	Id       int    `dbq:"id"`
	Kategori string `dbq:"kategori"`
	Nama     string `dbq:"nama"`
	Harga    int64  `dbq:"harga"`
	Jumlah   int    `dbq:"jumlah"`
}

func GetTableItem() string {
	return "item"
}

func TableItem() []string {
	return []string{
		"id",
		"kategori",
		"nama",
		"harga",
		"jumlah",
	}
}
