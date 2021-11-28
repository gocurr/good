package main

import (
	"fmt"
	"github.com/gocurr/good"
	log "github.com/sirupsen/logrus"
	"net/http"
)

var nameFns = good.NameFns{
	{"demo1", func() {
		db := good.DB()
		fmt.Println(db)

		producer := good.RocketMQProducer()
		fmt.Println(producer)

		client := good.TableStoreClient()
		fmt.Println(client)
		fmt.Println("demo1...")
	}},

	{"demo2", func() {
		fmt.Println("demo2...")
	}},
}

func main() {
	c := good.Configure("./app.yml", false)
	if err := good.StartCrontab(nameFns); err != nil {
		log.Fatalln(err)
	}

	good.ServerMux(http.NewServeMux())
	good.Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	good.Fire(c)
}
