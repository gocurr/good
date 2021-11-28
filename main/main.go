package main

import (
	"fmt"
	"github.com/gocurr/good"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func demo1() {
	log.Info("demo1...")
}

func demo2() {
	log.Info("demo2...")
}

func main() {
	good.RegisterCron("demo1", demo1)
	good.RegisterCron("demo2", demo2)
	if err := good.StartCrontab(); err != nil {
		log.Fatalln(err)
	}

	// good.ServerMux(http.NewServeMux())
	good.Route("/", func(w http.ResponseWriter, r *http.Request) {
		urls := good.Custom("urls")
		s, ok := urls.([]interface{})
		if ok {
			for _, a := range s {
				fmt.Println(a.(string))
			}
		}
		_, _ = w.Write([]byte("ok"))
	})

	good.Fire()
}
