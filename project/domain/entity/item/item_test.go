package item_test

import (
	"fmt"
	"game-store-final-project/project/domain/entity/item"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Positif Case
*/
func TestNewItem(t *testing.T) {
	item, err := item.NewItem(item.DTOItem{
		Id:       1,
		Nama:     "Xbox",
		Kategori: "New Console",
		Harga:    5432100,
		Jumlah:   1,
	})
	fmt.Println(item)
	assert.Nil(t, err)
	assert.Equal(t, "Xbox", item.GetNama())
}

/*
Negative Case
*/
func TestNewItemGetNama(t *testing.T) {
	_, err := item.NewItem(item.DTOItem{
		Id:       1,
		Nama:     "",
		Kategori: "New Console",
		Harga:    5432100,
		Jumlah:   1,
	})
	assert.NotNil(t, err)
	assert.Equal(t, "NAMA NOT SET", err.Error())
}

func TestNewItemGetKategori(t *testing.T) {
	_, err := item.NewItem(item.DTOItem{
		Id:       1,
		Nama:     "Xbox",
		Kategori: "",
		Harga:    5432100,
		Jumlah:   1,
	})
	assert.NotNil(t, err)
	assert.Equal(t, "CATEGORY ID NOT SET", err.Error())
}
