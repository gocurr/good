package main

import (
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
	c, _ := conf.NewDefault()

	_ = logger.Set(c)

	crons := crontab.New(c)
	_ = crons.Bind("demo1", func() {
		log.Info("demo1")
	})
	crons.Start()

	sugar.Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	sugar.Fire(c)
}
