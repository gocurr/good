package main

import (
	"github.com/gocurr/good"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	good.BindCron("demo1", func() {
		log.Info("demo1")
	})
	good.Register("hello", "*/5 * * * * ?", func() {
		log.Info("hello")
	})
	good.StartCrontab()

	good.ServerMux(http.NewServeMux())
	good.Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("good"))
	})
	good.Fire()
}
