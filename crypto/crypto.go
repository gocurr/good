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

const encLen = len("enc(")

// Decrypt returns a decrypted string by the given key and text.
func Decrypt(key, text string) (string, error) {
	if !isEnc(text) {
		return text, nil
	}

	enc := text[encLen : len(text)-1]
	return decrypt(key, enc)
}

// isEnc reports whether text is encrypted.
func isEnc(text string) bool {
	if len(text) <= len("enc()") {
		return false
	}

	prefix := text[:encLen]
	return (prefix == "enc(" || prefix == "ENC(") &&
		strings.HasSuffix(text, ")")
}

// decrypt returns the decrypted string by the given key and enc.
func decrypt(key, enc string) (string, error) {
	keyBytes, _ := hex.DecodeString(key)
	ciphertext, _ := hex.DecodeString(enc)

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("crypto: ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)
	return fmt.Sprintf("%s", ciphertext), nil
}

// Encrypt returns an encrypted string by the given key and plain.
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
