package good

import (
	"fmt"
	"net/http"
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

var secret = "9253a3c25e69cd7e469877b0c6005604"

func TestEnPw(t *testing.T) {
	encrypter, err := Encrypter(secret, "12345")
	if err != nil {
		panic(err)
	}
	fmt.Println(encrypter)
}

func TestDePw(t *testing.T) {
	decrypter, err := Decrypter(secret, "f80a78c6688ba43430e628539a4c6445b6ed5d9bf3")
	if err != nil {
		panic(err)
	}
	fmt.Println(decrypter)
}

func TestCreateSecret(t *testing.T) {
	secret = CreateSecret()
	fmt.Println(secret)
}

func TestPort(t *testing.T) {
	fmt.Println(1<<16 - 1)
}

func TestConfigure(t *testing.T) {
	Configure("./app.yml", false)
	if err := StartCrontab(nameFns); err != nil {
		panic(err)
	}

	ServerMux(http.NewServeMux())
	Route("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})
	Fire()
}
