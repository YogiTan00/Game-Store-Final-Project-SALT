package redis_test

//
//import (
//	"context"
//	"encoding/json"
//	"game-store-final-project/project/domain/entity/item"
//	"game-store-final-project/project/internal/config/redis"
//	redis2 "game-store-final-project/project/internal/repository/redis"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestRepoItemRedis_StoreOrUpdateDataItem(t *testing.T) {
//	redisConn := redis.InitRedisClient()
//	defer redisConn.Close()
//	ctx := context.Background()
//	repoItem := redis2.NewRepoItemRedis(redisConn)
//
//	dataItem, errItem := item.NewItem(item.DTOItem{
//		Id:       0,
//		Nama:     "",
//		Kategori: "",
//		Harga:    0,
//		Jumlah:   0,
//	})
//
//	err := repoItem.StoreOrUpdateListDataItem(ctx, dataBuku)
//
//	assert.Nil(t, errBuku)
//	assert.Nil(t, err)
//	dataJson := []byte(`{"id":1,"nama":"Service Console Xbox Sereis X","kategori":"Service Console","harga":"4673000","jumlah":"999"}`)
//
//	listDTOItem := make(item.DTOItem, 0)
//
//	errJson := json.Unmarshal(dataJson, &listDTOItem)
//
//	listItem := make(*item.Item, 0)
//
//	for _, dto := range listDTOItem {
//		newItem, _ := item.NewItem(dto)
//		listItem = append(listItem, newItem)
//	}
//
//	err := repoItem.StoreOrUpdateListDataItem(ctx, listItem)
//
//	assert.Nil(t, errJson)
//	assert.Nil(t, err)
//}
//
//func TestRepoItemRedis_GetItemByID(t *testing.T) {
//	redisConn := redis.InitRedisClient()
//	defer redisConn.Close()
//	ctx := context.Background()
//	repoItem := redis2.NewRepoItemRedis(redisConn)
//
//	_, errItem := item.NewItem(item.DTOItem{
//		Id:       11,
//		Nama:     "God of War Ragnarok",
//		Kategori: "Buy New Game",
//		Harga:    1029000,
//		Jumlah:   99,
//	})
//
//	dataFromRedis, errGet := repoItem.GetItemByID(ctx, "11")
//
//	assert.Nil(t, errItem)
//	assert.Nil(t, errGet)
//	assert.NotNil(t, dataFromRedis)
//}
