package examples

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/rocketmq"
	"testing"
)

func Test_Rocket(t *testing.T) {
	c, err := conf.New("../app.yaml")
	Panic(err)

	_, err = rocketmq.NewProducer(c)
	Panic(err)
}
