package main

import (
	"github.com/gocurr/good/sugar"
	"testing"
)

func Test_Gen(t *testing.T) {
	sugar.GenPasswd("123456", "secret.txt", true)
}
