package good

import (
	"testing"
)

func Test_Gen(t *testing.T) {
	secrect := CreateSecret()
	a, _ := Encrypter(secrect, "你好 hello abc")
	println(Decrypter(secrect, a))
}
