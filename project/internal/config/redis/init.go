package redis

import (
	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	rdb *redis.Client
}

//type Redis interface {
//	GetItemByID(ctx context.Context, id string) (*item.Item, error)
//}
//
//func InitRedis() Redis {
//	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
//	if err != nil {
//		panic(err)
//	}
//	rdb := redis.NewClient(opt)
//	return &redis.Client{}
//}

func InitRedisClient() *redis.Client {
	var (
		opts *redis.Options
	)

	opts = &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	}

	client := redis.NewClient(opts)

	return client
}
