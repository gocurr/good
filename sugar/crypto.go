package sugar

import (
	"fmt"
	"github.com/gocurr/good/crypto"
	"io/ioutil"
	"os"
)

// GenPasswd generates an encrypted string via pw and filename
func GenPasswd(pw, filename string) {
	var secret string
	all, err := ioutil.ReadFile(filename)
	if err != nil {
		secret = crypto.CreateSecret()
		if err := ioutil.WriteFile(filename, []byte(secret), os.ModePerm); err != nil {
			panic(err)
		}
	} else {
		secret = string(all)
	}

	encrypted, err := crypto.Encrypt(secret, pw)
	if err != nil {
		panic(err)
	}
	fmt.Printf("secret: %s\n", secret)
	fmt.Printf("encrypted: %s\n", encrypted)
}
