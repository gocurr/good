package main

import (
	"fmt"
	"github.com/gocurr/good/conf"
	"github.com/gocurr/good/sugar"
	"testing"
)

func TestConf(t *testing.T) {
	c, _ := conf.New("../application.yml")
	ab, err := c.ReservedString("a_B")
	if err != nil {
		return
	}
	fmt.Println(ab)
}

func TestRead(t *testing.T) {
	err := conf.Read("../application.yml", &custom)
	if err != nil {
		t.Errorf("%v", err)
		return
	}

	sugar.Fire(&custom)
}
