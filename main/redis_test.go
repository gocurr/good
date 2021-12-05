package main

import (
	"context"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/redis"
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_Redis(t *testing.T) {
	c, _ := conf.New("../app.yaml")
	rdb, err := redis.New(c)
	Panic(err)

	var ctx = context.Background()
	result, _ := rdb.Get(ctx, "good").Result()
	log.Info(result)

	rdb.Set(ctx, "good", "better", 0)
	result, _ = rdb.Get(ctx, "good").Result()
	log.Info(result)
}
