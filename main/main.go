package main

import (
	"fmt"
	"github.com/gocurr/good/conf"
)

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	c, err := conf.ReadDefault()
	Panic(err)
	fmt.Println(c.Int("key"))
	fmt.Println(c.String("str"))
	fmt.Println(c.Slice("urls"))
	fmt.Println(c.Map("complex"))
}
