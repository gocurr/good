package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/server"
	log "github.com/sirupsen/logrus"
	"net/http"
	"testing"
	"time"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

type Custom struct {
	Server struct {
		Port int `yaml:"port,omitempty"`
	} `yaml:"server,omitempty"`

	Logrus struct {
		TimeFormat string `yaml:"time_format,omitempty"`
		TTYDiscard bool   `yaml:"tty_discard,omitempty"`
		Graylog    struct {
			Enable bool                   `yaml:"enable,omitempty"`
			Host   string                 `yaml:"host,omitempty"`
			Port   int                    `yaml:"port,omitempty"`
			Extra  map[string]interface{} `yaml:"extra,omitempty"`
		} `yaml:"graylog,omitempty"`
	} `yaml:"logrus,omitempty"`

	Mysql struct {
		Driver     string `yaml:"driver,omitempty"`
		User       string `yaml:"user,omitempty"`
		Password   string `yaml:"password,omitempty"`
		Datasource string `yaml:"datasource,omitempty"`
	} `yaml:"mysql,omitempty"`

	Redis struct {
		Host     string `yaml:"host,omitempty"`
		Port     int    `yaml:"port,omitempty"`
		Password string `yaml:"password,omitempty"`
		DB       int    `yaml:"db,omitempty"`
	} `yaml:"redis,omitempty"`

	Crontab map[string]string `yaml:"crontab,omitempty"`

	Secure struct {
		Key string `yaml:"key,omitempty"`
	} `yaml:"secure,omitempty"`

	Pg struct {
		Addr     string `yaml:"addr"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"pg,omitempty"`

	Maria struct {
		Addr     string `yaml:"addr"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"maria,omitempty"`

	Logic struct {
		Api   string   `yaml:"api"`
		Names []string `yaml:"names"`
	} `yaml:"logic,omitempty"`
}

type Msg struct {
	Text string `json:"text"`
}

func Test_Main(t *testing.T) {
	var c Custom
	err := conf.Read("../app.yaml", &c)
	Panic(err)
	err = logger.Set(&c)
	Panic(err)

	server.Route("/", func(w http.ResponseWriter, r *http.Request) {
		hi := server.Parameter("hi", r)
		log.Infof("%s", hi)
		url := "http://127.0.0.1:9091"
		raw, err := server.HttpGetRaw(url)
		if err != nil {
			return
		}
		var out Msg
		if err = server.PostJSON(url, nil, &out); err != nil {
			log.Errorf("%v", err)
		}
		out.Text = time.Now().Format(consts.DefaultTimeFormat)
		log.Infof("%v", out)
		_, _ = w.Write(raw)
	})
	// server.Fire(9090)
}
