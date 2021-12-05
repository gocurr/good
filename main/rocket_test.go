package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/rocketmq"
	"github.com/gocurr/good/tablestore"
	"testing"
)

func Test_Rocket(t *testing.T) {
	c, err := conf.Read("../application.yml")
	if err != nil {
		return
	}

	_, _ = rocketmq.NewProducer(c)
}

func Test_Tablestore(t *testing.T) {
	c, err := conf.Read("../application.yml")
	if err != nil {
		return
	}

	_, _ = tablestore.New(c)
}
