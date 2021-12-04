package main

import (
	"github.com/gocurr/good/sugar"
	"testing"
)

func Test_Encrypted(t *testing.T) {
	sugar.Encrypted("123456", "secret.txt", true)
}
