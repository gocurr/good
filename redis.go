package good

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var ctx = context.Background()
var rdb *redis.Client

// initRedis inits rdb
func initRedis() error {
	redisConf := conf.Redis
	if redisConf == nil {
		return nil
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     redisConf.Host + ":" + strconv.Itoa(redisConf.Port),
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})
	_, err := rdb.Ping(ctx).Result()
	return err
}

// RedisDb returns rdb
func RedisDb() *redis.Client {
	return rdb
}
