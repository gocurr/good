package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/rocketmq"
	"testing"
)

func Test_Rocket(t *testing.T) {
	c, err := conf.New("../app.yaml")
	if err != nil {
		return
	}

	_, err = rocketmq.NewProducer(c)
	if err != nil {
		panic(err)
	}
}
