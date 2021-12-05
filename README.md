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
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/server"
	"net/http"
)

func main() {
	c, _ := conf.NewDefault()
	_ = logger.Set(c)

	server.Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	server.Fire(c)
}
```

### Custom

```go
package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/server"
	"net/http"
)

type Custom struct {
	Server struct {
		Port int `yaml:"port,omitempty"`
	} `yaml:"server,omitempty"`

	Logrus struct {
		TimeFormat     string `yaml:"time_format,omitempty"`
		TTYDiscard bool   `yaml:"tty_discard,omitempty"`
		Graylog    struct {
			Enable bool                   `yaml:"enable,omitempty"`
			Host   string                 `yaml:"host,omitempty"`
			Port   int                    `yaml:"port,omitempty"`
			Extra  map[string]interface{} `yaml:"extra,omitempty"`
		} `yaml:"graylog,omitempty"`
	} `yaml:"logrus,omitempty"`
}

func main() {
	var c Custom
	_ = conf.ReadDefault(&c)
	_ = logger.Set(&c)

	server.Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	server.Fire(&c)
}
```