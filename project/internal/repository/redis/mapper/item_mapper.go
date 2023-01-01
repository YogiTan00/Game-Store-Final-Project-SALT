package mapper

//import (
//	"encoding/json"
//	"errors"
//	"game-store-final-project/project/domain/entity/item"
//	"strconv"
//)
//
//type AttributeItem struct {
//	Id       int    `json:"id"`
//	Nama     string `json:"nama"`
//	Kategori string `json:"kategori"`
//	Harga    int64  `json:"harga"`
//	Jumlah   int    `json:"jumlah"`
//}
//
//func MapSetItemToArrayString(data []*item.Item) []string {
//	var listItem []string
//	for i := 0; i < len(data); i++ {
//		attrItem := &AttributeItem{
//			Id:       data[i].GetID(),
//			Nama:     data[i].GetNama(),
//			Kategori: data[i].GetKategori(),
//			Harga:    data[i].GetHarga(),
//			Jumlah:   data[i].GetJumlah(),
//		}
//		attrJson, _ := json.Marshal(attrItem)
//
//		listItem = append(listItem, strconv.Itoa(i))
//		listItem = append(listItem, string(attrJson))
//	}
//	return listItem
//}
//
//func MapSetItemToString(data *item.Item) string {
//	attrItem := &AttributeItem{
//		Id:       data.GetID(),
//		Nama:     data.GetNama(),
//		Kategori: data.GetKategori(),
//		Harga:    data.GetHarga(),
//		Jumlah:   data.GetJumlah(),
//	}
//
//	attrJson, _ := json.Marshal(attrItem)
//
//	return string(attrJson)
//}
//
//func MapGetRedisField() string {
//	return "data_item"
//}
//
//func MapGetKeyValueRedis(data *item.Item) string {
//	return string(data.GetID())
//}
//
//func MapFromJsonStringToDomainItem(data string) (*item.Item, error) {
//	attributeItem := new(AttributeItem)
//	err := json.Unmarshal([]byte(data), attributeItem)
//	if err != nil {
//		return nil, err
//	}
//
//	item, err := item.NewItem(item.DTOItem{
//		Id:       attributeItem.Id,
//		Nama:     attributeItem.Nama,
//		Kategori: attributeItem.Kategori,
//		Harga:    attributeItem.Harga,
//		Jumlah:   attributeItem.Jumlah,
//	})
//	if err != nil {
//		return nil, errors.New("gagal mapping data redis")
//	}
//
//	return item, nil
//}
