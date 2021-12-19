package examples

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
	_ = crons.Register("hello", "*/3 * * * * ?", func() {
		log.Info("hello...")
	})
	if err := crons.Start(); err != nil {
		panic(err)
	}

	time.Sleep(10 * time.Second)
}
