package good

import (
	"fmt"
	"testing"
)

var nameFns = NameFns{
	{"demo1", func() {
		fmt.Println("demo1...")
	}},

	{"demo2", func() {
		fmt.Println("demo2...")
	}},
}

func TestConfigure(t *testing.T) {
	Configure("./app.yml", false)

	if err := StartCrontab(nameFns); err != nil {
		panic(err)
	}

	select {}
}
