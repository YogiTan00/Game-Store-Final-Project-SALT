package redis

//
//import (
//	"context"
//	"errors"
//	"game-store-final-project/project/domain/entity/item"
//	mapper2 "game-store-final-project/project/internal/repository/redis/mapper"
//	"github.com/go-redis/redis/v8"
//)
//
//type RepoItemRedis struct {
//	Conn *redis.Client
//}
//
//type GetItemByID []*item.Item
//
//func NewRepoItemRedis(Conn *redis.Client) *RepoItemRedis {
//	return &RepoItemRedis{Conn: Conn}
//}
//
//func (item *RepoItemRedis) GetItemByID(ctx context.Context, id string) (*item.Item, error) {
//	var (
//		data     string
//		checkErr error
//	)
//	data, checkErr = item.Conn.HGet(ctx, id, "data_item").Result()
//	if checkErr != nil && checkErr != redis.Nil && data == "" {
//		return nil, errors.New("Data tidak ditemukan")
//	}
//
//	dataItem, err := mapper2.MapFromJsonStringToDomainItem(data)
//	if err != nil {
//		return nil, err
//	}
//
//	return dataItem, nil
//}
//
//func (item RepoItemRedis) StoreOrUpdateDataItem(ctx context.Context, data *item.Item) error {
//	item, err := item.Conn.HSet(ctx, mapper2.MapGetKeyValueRedis(data), "data_item", mapper2.MapSetItemToString(data)).Result()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (item RepoItemRedis) StoreOrUpdateListDataItem(ctx context.Context, data []*item.Item) error {
//	listItem := mapper2.MapSetItemToArrayString(data)
//	_, err := item.Conn.HMSet(ctx, "all_item", listItem).Result()
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
