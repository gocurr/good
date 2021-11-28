package good

import (
	"fmt"
	"io/ioutil"
	"os"
)

func GenPasswd(pw string) string {
	var secret string
	filename := "secret.txt"
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		secret = CreateSecret()
		fmt.Println(secret)
		if err := ioutil.WriteFile(filename, []byte(secret), os.ModePerm); err != nil {
			panic(err)
		}
	} else {
		secret = string(bytes)
	}

	encrypter, err := Encrypter(secret, pw)
	if err != nil {
		panic(err)
	}
	return encrypter
}
