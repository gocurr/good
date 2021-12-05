# The Go App Boot Framework

`good` is a `Go` framework that makes developers write go applications much easier.

## Download and Install

```bash
go get -u github.com/gocurr/good
```

## Usage

### Default

```go
package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/sugar"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	c, _ := conf.NewDefault()

	_ = logger.Set(c)

	crons := crontab.New(c)
	_ = crons.Bind("demo1", func() {
		log.Info("demo1")
	})
	crons.Start()

	sugar.Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	sugar.Fire(c)
}
```

### Custom struct

```go
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

	Crontab map[string]string `yaml:"crontab,omitempty"`

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

func main() {
	var c Custom
	if err := conf.ReadDefault(&c); err != nil {
		Panic(err)
	}

	_ = logger.Set(&c)
	crons := crontab.New(&c)
	_ = crons.Bind("demo1", func() {
		log.Info("demo1")
	})
	crons.Start()

	sugar.Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	sugar.Fire(&c)
}
```