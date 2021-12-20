package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
)

// Decrypt decrypts e through the given key
func Decrypt(key, e string) (string, error) {
	if !isEnc(e) {
		return e, nil
	}

	enc := e[len("enc[") : len(e)-1]
	return decrypt(key, enc)
}

// isEnc reports e is encrypted
func isEnc(e string) bool {
	if e == "" {
		return false
	}

	lower := strings.ToLower(e)
	return len(lower) > len("enc()") &&
		strings.HasPrefix(lower, "enc(") &&
		strings.HasSuffix(lower, ")")
}

// decrypt returns the decrypted string
// through the given key and the encrypted e
func decrypt(key, e string) (string, error) {
	keyBytes, _ := hex.DecodeString(key)
	ciphertext, _ := hex.DecodeString(e)

	block, err := aes.NewCipher(keyBytes)
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

// Encrypt returns an encrypted string
// through the given key and the plain text
func Encrypt(key, plain string) (string, error) {
	keyBytes, _ := hex.DecodeString(key)
	plaintext := []byte(plain)

	block, err := aes.NewCipher(keyBytes)
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

	return fmt.Sprintf("%x", ciphertext), nil
}
