package good_test

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/httpclient"
	"github.com/gocurr/good/mysql"
	"github.com/gocurr/good/postgres"
	"github.com/gocurr/good/redis"
	"github.com/gocurr/good/rocketmq"
	"github.com/gocurr/good/tablestore"
	"testing"
	"time"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func Test_All(t *testing.T) {
	// configuration
	c, err := conf.NewDefault()
	Panic(err)

	// crontab
	crons, err := crontab.New(c)
	Panic(err)
	err = crons.Bind("demo1", func() {
		t.Log("demo1")
	})
	Panic(err)
	err = crons.Register("hello", "*/3 * * * * ?", func() {
		t.Log("hello...")
	})
	Panic(err)

	// Custom struct
	type Custom struct {
		Server struct {
			Port int `yaml:"port,omitempty"`
		} `yaml:"server,omitempty"`
	}
	var custom Custom
	err = conf.ReadDefault(&custom)
	Panic(err)

	// httpclient
	type msg struct {
		Text string
	}
	url := "http://127.0.0.1:9090"
	var out msg
	_ = httpclient.PostJSON(url, msg{Text: "great"}, &out, time.Second)
	_, _ = httpclient.GetRaw(url, time.Millisecond)

	// mysql
	_, err = mysql.Open(c)
	Panic(err)

	// redis
	_, err = redis.New(c, 1)
	Panic(err)

	// postgres
	_, err = postgres.Open(c)
	Panic(err)

	// rocketmq
	_, err = rocketmq.NewProducer(c)
	Panic(err)

	// tablestore
	_, err = tablestore.New(c)
	Panic(err)
}
