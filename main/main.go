package main

import (
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/sugar"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

type Custom struct {
	Pg struct {
		Addr     string `yaml:"addr"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"pg"`

	Maria struct {
		Addr     string `yaml:"addr"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"maria"`

	Logic struct {
		Api   string   `yaml:"api"`
		Names []string `yaml:"names"`
	}
}

var custom Custom

func main() {
	c, _ := conf.ReadDefault()
	_ = logger.Set(c)
	crons := crontab.New(c)
	_ = crons.Bind("demo1", func() {
		log.Info("demo1")
	})
	_ = crons.Bind("demo2", func() {
		log.Info("demo2")
	})
	crons.Register("demo2", "*/1 * * * * ?", func() {
		log.Info("demo3")
	})
	crons.Start()

	fmt.Println(c.Int("xxx"))
	fmt.Println(c.String("key"))

	if err := c.Fill(&custom); err != nil {
		log.Error(err)
	}
	fmt.Println(custom)

	mysqlOp(c)
	redisOp(c)

	type msg struct {
		Text string `json:"text"`
	}
	sugar.Route("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := sugar.PostJSON("http://127.0.0.1:9091", &msg{Text: "hello"})
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		_, _ = w.Write(bytes)
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("test"))
	})

	sugar.ServerMux(mux)
	sugar.Fire(c)
}
