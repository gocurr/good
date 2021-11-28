package good

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

var RedisContext = context.Background()
var RedisDB *redis.Client

func initRedis() error {
	redisConf := conf.Redis
	if redisConf == nil {
		return nil
	}

	RedisDB = redis.NewClient(&redis.Options{
		Addr:     redisConf.Host + ":" + strconv.Itoa(redisConf.Port),
		Password: redisConf.Password,
		DB:       redisConf.DB,
	})
	_, err := RedisDB.Ping(RedisContext).Result()
	return err
}
