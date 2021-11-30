package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/sugar"
	log "github.com/sirupsen/logrus"
	"net/http"
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
	crontab.Register("hello", "*/10 * * * * ?", func() {
		log.Infof("hello")
	})
	err = crontab.StartCrontab()
	panic_(err)

	/*err = mysql.Init(c)
	panic_(err)
	rows, err := mysql.DB.Query("select name from names")
	panic_(err)
	defer func() { _ = rows.Close() }()
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
	fmt.Println(urls)*/

	sugar.Route("/", func(w http.ResponseWriter, r *http.Request) {
		//sugar.JSONHeader(w)
		url := sugar.Parameter("url", r)
		log.Infof("key = %v", url)

		var data = []byte("ok")
		get, err := sugar.HttpGet(url)
		if err == nil {
			data = get
		}
		_, _ = w.Write(data)
	})
	sugar.Fire(c)
}
