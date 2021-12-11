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

func Test_Time(t *testing.T) {
	println(sugar.NowString("2006"))
	println(sugar.NowString("2006-01-02"))
	println(sugar.NowString())

	fmt.Printf("%v\n", sugar.ParseTime("2021-12-06", "2006-01-02"))
	fmt.Printf("%v\n", sugar.ParseTime("2021-12-06 20:40:12", consts.DefaultTimeFormat))

	println(sugar.FormatTime(time.Now(), "15:04"))

	println(sugar.PointTimeString("20211206", 1, time.Minute*15))
	fmt.Printf("%v\n", sugar.PointTime("20211206", 1, time.Minute*15))
}
