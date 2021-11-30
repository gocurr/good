package main

import (
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	"github.com/gocurr/good/sugar"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	c, _ := conf.ReadDefault()
	crontab.Init(c)
	_ = crontab.Bind("demo1", func() {
		log.Info("demo1")
	})
	_ = crontab.Bind("demo2", func() {
		log.Info("demo2")
	})
	//crontab.Start()

	fmt.Println(c.Int("xxx"))
	fmt.Println(c.String("key", false))
	fmt.Println(c.String("key", true))

	//mysqlOp(c)
	//redisOp(c)

	sugar.Route("/", func(w http.ResponseWriter, r *http.Request) {
		sugar.JSONHeader(w)
		sugar.HandleErr(w, http.StatusForbidden, nil)
	})
	sugar.Fire(c)
}
