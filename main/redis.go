package main

import (
	"context"
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/redis"
)

func redisOp(c *conf.Configuration) {
	rdb, err := redis.Get(c)
	Panic(err)

	var ctx = context.Background()
	result, err := rdb.Get(ctx, "a").Result()
	Panic(err)
	fmt.Println(result)

	rdb.Set(ctx, "a", "nice", 0)
	result, err = rdb.Get(ctx, "a").Result()
	Panic(err)
	fmt.Println(result)
}
