package good

import (
	"fmt"
	"github.com/gocurr/good/crypto"
	"io/ioutil"
	"os"
)

func GenPasswd(pw string) string {
	var secret string
	filename := "secret.txt"
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		secret = crypto.CreateSecret()
		fmt.Println(secret)
		if err := ioutil.WriteFile(filename, []byte(secret), os.ModePerm); err != nil {
			panic(err)
		}
	} else {
		secret = string(bytes)
	}

	encrypter, err := crypto.Encrypter(secret, pw)
	if err != nil {
		panic(err)
	}
	return encrypter
}
