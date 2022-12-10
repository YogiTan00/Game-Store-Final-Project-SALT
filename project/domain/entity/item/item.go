package item

import "errors"

type Item struct {
	id       int
	nama     string
	kategori string
	harga    int64
	jumlah   int
}

type DTOItem struct {
	Id       int
	Nama     string
	Kategori string
	Harga    int64
	Jumlah   int
}

func NewItem(t DTOItem) (*Item, error) {
	if t.Nama == "" {
		return nil, errors.New("NAMA NOT SET")
	}
	if t.Kategori == "" {
		return nil, errors.New("CATEGORY ID NOT SET")
	}

	return &Item{
		id:       t.Id,
		kategori: t.Kategori,
		nama:     t.Nama,
		harga:    t.Harga,
		jumlah:   t.Jumlah,
	}, nil
}

func NewListItem(t []DTOItem) ([]*Item, error) {
	listItem := make([]*Item, 0)
	for _, data := range t {
		item, err := NewItem(data)
		if err != nil {
			return nil, err
		}
		listItem = append(listItem, item)
	}

	return listItem, nil
}

func (i *Item) GetID() int {
	return i.id
}

func (i *Item) GetNama() string {
	return i.nama
}

func (i *Item) GetKategori() string {
	return i.kategori
}

func (i *Item) GetHarga() int64 {
	return i.harga
}

func (i *Item) GetJumlah() int {
	return i.jumlah
}
