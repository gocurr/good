package main

import (
	"github.com/gocurr/good"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func demo1() {
	log.Info("demo1...")
}

func demo2() {
	log.Info("demo2...")
}

func main() {
	// good.Configure("app.yml", false)
	good.ConfigDefault()

	good.RegisterCron("demo1", demo1)
	good.RegisterCron("demo2", demo2)
	if err := good.StartCrontab(); err != nil {
		log.Fatalln(err)
	}

	good.Route("/", func(w http.ResponseWriter, r *http.Request) {
		good.JsonHeader(w)
		urls := good.Custom("urls")
		s, ok := urls.([]interface{})
		if ok {
			for _, a := range s {
				log.Info(a.(string))
			}
		}

		good.RegisterCron("a", func() {

		})
		if err := good.StartCrontab(); err != nil {
			log.Error(err)
		}

		key := good.Parameters("key", r)
		if key != nil {
			log.Info(key)
		}

		p := good.Parameter("good", r)
		log.Info(p)

		bytes, err := good.JSONBytes(r)
		if err == nil {
			log.Info(string(bytes))
		}

		println(time.Now().Format(good.DefaultTimeFormat))
		_, _ = w.Write([]byte(`{"data":"ok"}`))
	})

	//good.ServerMux(http.NewServeMux())
	good.Fire(callback)
}

func callback() {
	log.Info("hello app")
}
