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

func TestRoundFloat(t *testing.T) {
	fmt.Printf("%v", RoundFloat(1.015, 2))
}

func TestCeilFloat(t *testing.T) {
	fmt.Printf("%v", CeilFloat(1.21, 2))
}

func TestFloorFloat(t *testing.T) {
	fmt.Printf("%v", FloorFloat(1.21, 2))
}

func TestPostJSON(t *testing.T) {
	var data []interface{}
	json, err := PostJSON("http://127.0.0.1:9090", data)
	if err != nil {
		return
	}
	fmt.Printf("%s", string(json))
}

func TestHttpGet(t *testing.T) {
	get, err := HttpGet("http://127.0.0.1:9090")
	if err != nil {
		return
	}
	fmt.Printf("%s", string(get))
}
