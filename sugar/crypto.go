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

var hexes = []rune("0123456789abcdefABCDEF")

// newKey returns a random 32bits cipher-key.
func newKey() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	var b strings.Builder
	for i := 0; i < 32; i++ {
		a := r.Intn(len(hexes))
		b.WriteRune(hexes[a])
	}
	return b.String()
}

// PrintKeyEnc prints cipher-key and encrypted text.
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
