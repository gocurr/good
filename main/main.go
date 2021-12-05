package main

import (
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
	Server struct {
		Port int `yaml:"port"`
	}

	Logrus struct {
		Format  string `yaml:"format,omitempty"`
		TTY     bool   `yaml:"tty,omitempty"`
		GrayLog struct {
			Enable bool                   `yaml:"enable,omitempty"`
			Host   string                 `yaml:"host,omitempty"`
			Port   int                    `yaml:"port,omitempty"`
			Extra  map[string]interface{} `yaml:"extra,omitempty"`
		} `yaml:"graylog,omitempty"`
	}

	Mysql struct {
		Driver     string `yaml:"driver,omitempty"`
		User       string `yaml:"user,omitempty"`
		Password   string `yaml:"password,omitempty"`
		Datasource string `yaml:"datasource,omitempty"`
	}

	Redis struct {
		Host     string `yaml:"host,omitempty"`
		Port     int    `yaml:"port,omitempty"`
		Password string `yaml:"password,omitempty"`
		DB       int    `yaml:"db,omitempty"`
	}

	Crontab map[string]string

	Secure struct {
		Key string `yaml:"key,omitempty"`
	}

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

var c Custom

func main() {
	//cc, err := conf.NewDefault(&custom)
	//if err == nil {
	//	sugar.Fire(cc)
	//	//sugar.Fire(custom)
	//	return
	//}

	//c, _ := conf.NewDefault()
	err := conf.ReadDefault(&c)
	if err != nil {
		return
	}
	_ = logger.Set(c)
	crons := crontab.New(c)
	_ = crons.Bind("demo1", func() {
		log.Info("demo1")
	})
	_ = crons.Bind("demo2", func() {
		log.Info("demo2")
	})
	crons.Register("demo3", "*/1 * * * * ?", func() {
		log.Info("demo3")
	})
	crons.Start()

	mysqlOp(c)
	redisOp(c)

	type msg struct {
		Text string `json:"text"`
	}
	sugar.Route("/", func(w http.ResponseWriter, r *http.Request) {
		var out interface{}
		err := sugar.PostJSON("http://127.0.0.1:9091", &msg{Text: "hello"}, &out)
		if err != nil {
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		log.Info(out)
		_, _ = w.Write([]byte("ok"))
	})

	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("test"))
	})

	sugar.ServerMux(mux)
	sugar.Fire(c)
}
