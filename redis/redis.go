package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crypto"
	"strconv"
)

var redisErr = errors.New("bad redis configuration")

// Get returns *redis.Client
func Get(c *conf.Configuration) (*redis.Client, error) {
	redisConf := c.Redis
	if redisConf == nil {
		return nil, redisErr
	}

	var err error
	var pw string
	if c.Secure == nil || c.Secure.Key == "" {
		pw = redisConf.Password
	} else {
		pw, err = crypto.Decrypt(c.Secure.Key, redisConf.Password)
		if err != nil {
			return nil, err
		}
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConf.Host + ":" + strconv.Itoa(redisConf.Port),
		Password: pw,
		DB:       redisConf.DB,
	})
	_, err = rdb.Ping(context.Background()).Result()
	return rdb, err
}
