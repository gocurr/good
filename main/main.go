package main

import (
	"github.com/gocurr/good"
	log "github.com/sirupsen/logrus"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	good.BindCron("demo1", func() {
		log.Info("demo1")
	})
	good.BindCron("demo2", func() {
		log.Info("demo2")
	})
	good.StartCrontab()
	good.Fire()
}
