package main

import (
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/logger"
	"github.com/gocurr/good/sugar"
	"net/http"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	c, err := conf.ReadDefault()
	Panic(err)
	err = logger.Init(c)
	Panic(err)
	sugar.Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	sugar.Fire(c)
}
