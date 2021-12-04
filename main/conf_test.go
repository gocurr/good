package main

import (
	"fmt"
	"github.com/gocurr/good/conf"
	"testing"
)

func TestConf(t *testing.T) {
	c, _ := conf.Read("../application.yml")
	ab, err := c.String("a_B")
	if err != nil {
		return
	}
	fmt.Println(ab)
}
