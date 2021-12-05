package main

import (
	"fmt"
	"github.com/gocurr/good/consts"
	"github.com/gocurr/good/sugar"
	"testing"
	"time"
)

func Test_Encrypted(t *testing.T) {
	sugar.PrintKeyEnc("123456", "secret.txt")
}

func Test_TimeFormat(t *testing.T) {
	println(time.Now().Format(consts.DefaultTimeFormat))
}

func Test_Float(t *testing.T) {
	fmt.Println(sugar.RoundFloat(1.2345, 2))
	fmt.Println(sugar.CeilFloat(1.2345, 2))
	fmt.Println(sugar.FloorFloat(1.2345, 2))

	fmt.Println(sugar.RoundFloat(1.2345, 3))
	fmt.Println(sugar.CeilFloat(1.2345, 3))
	fmt.Println(sugar.FloorFloat(1.2345, 3))

	quotient := sugar.FloatQuotient(1, 3)
	fmt.Println(quotient)
	fmt.Println(sugar.RoundFloat(quotient, 2))
}
