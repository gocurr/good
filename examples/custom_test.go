package examples

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/server"
	"net/http"
	"testing"
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
		TimeFormat string `yaml:"time-format,omitempty"`
		TTYDiscard bool   `yaml:"tty-discard,omitempty"`
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

	Crontab struct {
		Enable bool              `yaml:"enable,omitempty"`
		Specs  map[string]string `yaml:"specs,omitempty"`
	} `yaml:"crontab,omitempty"`

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

func Test_Custom(t *testing.T) {
	var c Custom
	err := conf.Read("../app.yaml", &c)
	Panic(err)

	err = logger.Set(&c)
	Panic(err)

	server.Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello better"))
	})
	// server.Fire(9090)
}
