package main

import (
	"fmt"
	"github.com/gocurr/good"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var nameFns = good.NameFns{
	{"demo1", func() {
		redis, ctx := good.Redis()
		result, _ := redis.Get(ctx, "a").Result()
		fmt.Println(result)
	}},
	{"demo2", func() {
		redis, ctx := good.Redis()
		result, _ := redis.Get(ctx, "b").Result()
		fmt.Println(result)
	}},
}

func main() {
	good.Configure("./app.yml", false)
	if err := good.StartCrontab(nameFns); err != nil {
		log.Errorf("%v", err)
	}

	// good.ServerMux(http.NewServeMux())
	good.Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	good.Fire()
}
