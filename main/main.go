package main

import (
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/crontab"
	log "github.com/sirupsen/logrus"
	"time"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	c, err := conf.ReadDefault()
	Panic(err)
	fmt.Println(c.String("str"))
	fmt.Println(c.Int("key"))
	slice := c.Slice("urls")
	fmt.Println(slice)
	m := c.Map("complex")
	for k, v := range m {
		fmt.Println(k, v)
	}
	crontab.Register("hello", "*/1 * * * * ?", func() {
		log.Info("hello")
	})
	err = crontab.Start()
	Panic(err)
	time.Sleep(1 * time.Minute)
}
