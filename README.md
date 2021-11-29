# The Go App Boot Framework

`good` is a http framework that makes developers write go applications much easier.

## Download and Install

```bash
go get -u github.com/gocurr/good
```

## Usage

### Custom

```go
package main

import (
	"context"
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/db"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/redis"
	"github.com/gocurr/good/rocketmq"
	"github.com/gocurr/good/tablestore"
	log "github.com/sirupsen/logrus"
	"time"
)

func panic_(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	c, err := conf.Read("application.yml")
	panic_(err)

	err = logger.Init(c)
	panic_(err)

	crontab.Init(c)
	crontab.Register("hello", "*/1 * * * * ?", func() {
		log.Infof("hello")
	})
	err = crontab.StartCrontab()
	panic_(err)

	err = db.Init(c)
	panic_(err)
	rows, err := db.DB.Query("select abc from abc")
	panic_(err)
	for rows.Next() {
		var abc string
		err = rows.Scan(&abc)
		panic_(err)
		log.Infof("from db: %v", abc)
	}

	err = tablestore.Init(c)
	panic_(err)

	err = rocketmq.Init(c)
	panic_(err)

	err = redis.Init(c)
	panic_(err)
	result, err := redis.Rdb.Get(context.Background(), "a").Result()
	panic_(err)
	log.Info("redis------", result)

	urls := c.Custom["urls"].([]interface{})
	fmt.Println(urls)

	time.Sleep(1 * time.Minute)
}
```

### Simple

```go
package main

import (
	"github.com/gocurr/good"
	"net/http"
)

func main() {
	// good.ServerMux(http.NewServeMux())
	good.Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	good.Fire()
}
```