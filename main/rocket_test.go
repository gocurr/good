package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/rocketmq"
	"github.com/gocurr/good/tablestore"
	log "github.com/sirupsen/logrus"
	"testing"
)

func Test_Rocket(t *testing.T) {
	c, err := conf.New("../app.yaml")
	if err != nil {
		return
	}

	_, _ = rocketmq.NewProducer(c)
}

func Test_Tablestore(t *testing.T) {
	c, err := conf.New("../app.yaml")
	if err != nil {
		log.Error(err)
		return
	}

	_, err = tablestore.New(c)
	if err != nil {
		log.Error(err)
	}
}
