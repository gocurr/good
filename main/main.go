package main

import (
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/sugar"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	c, _ := conf.ReadDefault()
	_ = logger.Init(c)
	crons := crontab.New(c)
	_ = crons.Bind("demo1", func() {
		log.Info("demo1")
	})
	_ = crons.Bind("demo2", func() {
		log.Info("demo2")
	})
	crons.Register("demo2", "*/1 * * * * ?", func() {
		log.Info("demo3")
	})
	crons.Start()

	fmt.Println(c.Int("xxx"))
	fmt.Println(c.String("key", false))
	fmt.Println(c.String("key", true))

	mysqlOp(c)
	redisOp(c)

	sugar.Route("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := sugar.HttpGet("http://127.0.0.1:9091")
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		_, _ = w.Write(bytes)
	})
	sugar.Fire(c)
}
