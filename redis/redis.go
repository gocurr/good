package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
	"strconv"
)

// Rdb the global redis client
var Rdb *redis.Client

// Init inits rdb
func Init(c *conf.Configuration) error {
	redisConf := c.Redis
	pw, err := crypto.Decrypt(c.Secure.Key, redisConf.Password)
	if err != nil {
		return err
	}
	Rdb = redis.NewClient(&redis.Options{
		Addr:     redisConf.Host + ":" + strconv.Itoa(redisConf.Port),
		Password: pw,
		DB:       redisConf.DB,
	})
	_, err = Rdb.Ping(context.Background()).Result()
	return err
}
