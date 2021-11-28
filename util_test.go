package good

import (
	"fmt"
	"github.com/gocurr/good/crypto"
	"testing"
)

func TestGenPasswd(t *testing.T) {
	println(GenPasswd("hello"))
}

func TestCrypto(t *testing.T) {
	decrypt, err := crypto.Decrypt("73849a30e39de97b402d67085b20e04b", "enc(53bb3884c96203bc295002eadfac764f451dd1b978)")
	if err != nil {
		return
	}
	fmt.Println(decrypt)
}
