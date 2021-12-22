package examples

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/logger"
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_Crontab(t *testing.T) {
	c, err := conf.New("../app.yaml")
	Panic(err)
	err = logger.Set(c)
	Panic(err)

	crons, err := crontab.New(c)
	Panic(err)
	err = crons.Bind("demo1", func() {
		log.Info("demo1")
	})
	Panic(err)
	err = crons.Register("hello", "*/3 * * * * ?", func() {
		log.Info("hello...")
	})
	Panic(err)
	//err = crons.Start()
	//Panic(err)

	//time.Sleep(10 * time.Second)
}
