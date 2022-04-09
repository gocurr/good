# Go boot layer of frameworks for developers

`good` is a `Go` boot layer of frameworks that makes developers write applications much easier.

## Download and Install

```bash
go get -u github.com/gocurr/good
```

## Usage

#### Use environment variable to set secure key
```bash
export GOOD_SECURE_KEY=6c841a6c8b9e3b42dfd30dc950da0382
```
Or

### Use yaml configuration to set secure key
```yaml
secure:
  key: 6c841a6c8b9e3b42dfd30dc950da0382
```

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
		TimeFormat string `yaml:"time-format,omitempty"`
		TTYDiscard bool   `yaml:"tty-discard,omitempty"`
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