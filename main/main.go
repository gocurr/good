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
	fmt.Println(c.String("str"))
	fmt.Println(c.Int("key"))
	slice := c.Slice("urls")
	fmt.Println(slice)
	m := c.Map("complex")
	for k, v := range m {
		fmt.Println(k, v)
	}
}
