package mapper

import (
	"game-store-final-project/project/domain/entity/item"
	"game-store-final-project/project/internal/repository/mysql/model"
)

func DataDbToEntityItem(dataDTO item.DTOItem) (*item.Item, error) {
	post, err := item.NewItem(dataDTO)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func DataListDbToEntityItem(dataDTO []item.DTOItem) ([]*item.Item, error) {
	listItem := make([]*item.Item, 0)
	for _, data := range dataDTO {
		item, err := item.NewItem(data)
		if err != nil {
			return nil, err
		}
		listItem = append(listItem, item)
	}

	return listItem, nil
}

func ModelToDomainItem(m *model.ItemModel) (*item.Item, error) {
	posting, err := item.NewItem(item.DTOItem{
		Id:       m.Id,
		Kategori: m.Kategori,
		Nama:     m.Nama,
		Harga:    m.Harga,
		Jumlah:   m.Jumlah,
	})
	if err != nil {
		return nil, err
	}

	return posting, nil
}

func ListModelToDomainItem(m []*model.ItemModel) ([]*item.Item, error) {
	list := make([]*item.Item, 0)
	for _, modelData := range m {
		d, _ := ModelToDomainItem(modelData)
		list = append(list, d)
	}
	return list, nil
}
