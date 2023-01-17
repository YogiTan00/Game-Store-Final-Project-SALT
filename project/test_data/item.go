package test_data

import (
	"game-store-final-project/project/domain/entity/item"
)

func GetTestDataItem() *item.Item {
	item, _ := item.NewItem(item.DTOItem{
		Id:       1,
		Nama:     "Xbox",
		Kategori: "New Console",
		Harga:    5432100,
		Jumlah:   1,
	})
	return item
}

func GetTestDataCountItem(count int) []*item.Item {
	listItem := make([]*item.Item, 0)
	for i := 1; i <= count; i++ {
		item, _ := item.NewItem(item.DTOItem{
			Id:       1,
			Nama:     "Xbox",
			Kategori: "New Console",
			Harga:    5432100,
			Jumlah:   1,
		})

		listItem = append(listItem, item)
	}
	return listItem
}
