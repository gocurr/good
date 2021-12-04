package main

import (
	"github.com/gocurr/good/sugar"
	"testing"
)

func Test_Encrypted(t *testing.T) {
	sugar.PrintKeyEnc("123456", "secret.txt")
}
