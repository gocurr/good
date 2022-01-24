package good_test

import (
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/grpc"
	"github.com/gocurr/good/httpclient"
	"github.com/gocurr/good/mysql"
	"github.com/gocurr/good/postgres"
	"github.com/gocurr/good/redis"
	"github.com/gocurr/good/rocketmq"
	"github.com/gocurr/good/tablestore"
	"testing"
)

func handleErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func Test_All(t *testing.T) {
	// configuration
	c, err := conf.NewDefault()
	handleErr(err)

	ok, addr, timeout := grpc.ClientAddrTimeout(c)
	if ok {
		fmt.Println(addr, timeout)
	}

	ok, port := grpc.ServerPort(c)
	if ok {
		fmt.Println(port)
	}

	// crontab
	crons, err := crontab.New(c)
	handleErr(err)
	err = crons.Bind("demo1", func() {
		t.Log("demo1")
	})
	handleErr(err)
	err = crons.Register("hello", "*/3 * * * * ?", func() {
		t.Log("hello...")
	})
	handleErr(err)

	// Custom struct
	type Custom struct {
		Server struct {
			Port int `yaml:"port,omitempty"`
		} `yaml:"server,omitempty"`
	}
	var custom Custom
	err = conf.ReadDefault(&custom)
	handleErr(err)

	// httpclient
	type msg struct {
		Text string
	}
	url := "http://127.0.0.1:9090"
	var out msg
	_ = httpclient.PostJSON(url, msg{Text: "great"}, &out)
	_, _ = httpclient.GetRaw(url)

	// mysql
	_, err = mysql.Open(c)
	handleErr(err)

	// redis
	_, err = redis.New(c)
	handleErr(err)

	// postgres
	_, err = postgres.Open(c)
	handleErr(err)

	// rocketmq
	_, err = rocketmq.NewProducer(c)
	handleErr(err)

	// tablestore
	_, err = tablestore.New(c)
	handleErr(err)
}
