package main

import (
	"context"
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/db"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/redis"
	"github.com/gocurr/good/rocketmq"
	"github.com/gocurr/good/tablestore"
	log "github.com/sirupsen/logrus"
	"time"
)

func panic_(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	c, err := conf.Read("application.yml")
	panic_(err)

	err = logger.Init(c)
	panic_(err)

	crontab.Init(c)
	crontab.Register("hello", "*/1 * * * * ?", func() {
		log.Infof("hello")
	})
	err = crontab.StartCrontab()
	panic_(err)

	err = db.Init(c)
	panic_(err)
	rows, err := db.DB.Query("select name from names")
	panic_(err)
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		panic_(err)
		log.Infof("database name: %v", name)
	}

	err = tablestore.Init(c)
	panic_(err)

	err = rocketmq.Init(c)
	panic_(err)

	err = redis.Init(c)
	panic_(err)
	result, err := redis.Rdb.Get(context.Background(), "a").Result()
	panic_(err)
	log.Info("redis------", result)

	urls := c.Custom["urls"].([]interface{})
	fmt.Println(urls)

	time.Sleep(1 * time.Minute)
}
