package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	mrand "math/rand"
	"strings"
	"time"
)

var hexs = []rune{
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f',
}

// Decrypt decrypts str
func Decrypt(key, str string) (string, error) {
	if !isEnc(str) {
		return str, nil
	}

	enc := str[len("enc[") : len(str)-1]
	return decrypt(key, enc)
}

// isEnc reports str is encrypted
func isEnc(str string) bool {
	if str == "" {
		return false
	}

	lower := strings.ToLower(str)
	return len(lower) > len("enc()") &&
		strings.HasPrefix(lower, "enc(") &&
		strings.HasSuffix(lower, ")")
}

// CreateSecret returns a secret key
func CreateSecret() string {
	mrand.Seed(time.Now().UnixNano())
	var builder strings.Builder
	for i := 0; i < len("6368616e676520746869732070617373"); i++ {
		a := mrand.Intn(len(hexs) - 1)
		builder.WriteRune(hexs[a])
	}
	return builder.String()
}

// decrypt text with secret
func decrypt(secret, text string) (string, error) {
	key, _ := hex.DecodeString(secret)
	ciphertext, _ := hex.DecodeString(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext), nil
}

// Encrypt msg with secret
func Encrypt(secret, msg string) (string, error) {
	key, _ := hex.DecodeString(secret)
	plaintext := []byte(msg)

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return fmt.Sprintf("%x\n", ciphertext), nil
}
