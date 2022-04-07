package main

import (
	"context"
	"fmt"

	"github.com/gocurr/good/conf"
	_ "github.com/gocurr/good/crontab"
	_ "github.com/gocurr/good/crypto"
	_ "github.com/gocurr/good/httpclient"
	_ "github.com/gocurr/good/logger"
	_ "github.com/gocurr/good/mysql"
	_ "github.com/gocurr/good/oracle"
	_ "github.com/gocurr/good/postgres"
	"github.com/gocurr/good/redis"
	_ "github.com/gocurr/good/rocketmq"
	_ "github.com/gocurr/good/server"
	_ "github.com/gocurr/good/sugar"
	_ "github.com/gocurr/good/tablestore"
)

// For pkg-updates only.
func main() {
	c, err := conf.New("../app.yaml")
	if err != nil {
		panic(err)
	}

	rdb, err := redis.New(c)
	if err != nil {
		panic(err)
	}
	s, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
