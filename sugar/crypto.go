package sugar

import (
	"fmt"
	"github.com/gocurr/good/crypto"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var hexes = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f',
}

// Key returns a secret key
func Key() string {
	rand.Seed(time.Now().UnixNano())
	var builder strings.Builder
	for i := 0; i < 32; i++ {
		a := rand.Intn(len(hexes))
		builder.WriteRune(hexes[a])
	}
	return builder.String()
}

// PrintKeyEnc prints secret-key and encrypted string
func PrintKeyEnc(pw, filename string, reset ...bool) {
	var secret string
	if len(reset) > 0 && reset[0] {
		_ = os.Remove(filename)
	}

	all, err := ioutil.ReadFile(filename)
	if err != nil {
		secret = Key()
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
