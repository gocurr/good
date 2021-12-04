package main

import (
	"errors"
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
	cron := crontab.New(c)
	_ = cron.Bind("demo1", func() {
		log.Info("demo1")
	})
	_ = cron.Bind("demo2", func() {
		log.Info("demo2")
	})
	cron.Start()

	fmt.Println(c.Int("xxx"))
	fmt.Println(c.String("key", false))
	fmt.Println(c.String("key", true))

	mysqlOp(c)
	redisOp(c)

	sugar.Route("/", func(w http.ResponseWriter, r *http.Request) {
		sugar.JSONHeader(w)
		sugar.HandleErr(errors.New("some err"), w)
	})
	sugar.Fire(c)
}
