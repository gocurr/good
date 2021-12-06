package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/logger"
	log "github.com/sirupsen/logrus"
	"testing"
	"time"
)

func Test_Crontab(t *testing.T) {
	c, _ := conf.New("../app.yaml")
	_ = logger.Set(c)

	crons, err := crontab.New(c)
	if err != nil {
		Panic(err)
	}
	_ = crons.Bind("demo1", func() {
		log.Info("demo1")
	})
	crons.Register("hello", "*/3 * * * * ?", func() {
		time.Sleep(9 * time.Second)
		log.Info("hello...")
	})
	crons.Start()

	time.Sleep(10 * time.Second)
}
