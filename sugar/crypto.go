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

// newKey returns a secret newKey
func newKey() string {
	rand.Seed(time.Now().UnixNano())
	var builder strings.Builder
	for i := 0; i < 32; i++ {
		a := rand.Intn(len(hexes))
		builder.WriteRune(hexes[a])
	}
	return builder.String()
}

// PrintKeyEnc prints secret-newKey and encrypted string
func PrintKeyEnc(pw, filename string, reset ...bool) {
	var key string
	if len(reset) > 0 && reset[0] {
		_ = os.Remove(filename)
	}

	all, err := ioutil.ReadFile(filename)
	if err != nil {
		key = newKey()
		if err := ioutil.WriteFile(filename, []byte(key), os.ModePerm); err != nil {
			panic(err)
		}
	} else {
		key = string(all)
	}

	encrypted, err := crypto.Encrypt(key, pw)
	if err != nil {
		panic(err)
	}
	fmt.Printf("key: %s\n", key)
	fmt.Printf("encrypted: %s\n", encrypted)
}
