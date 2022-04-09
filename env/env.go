package env

import (
	"os"
)

// GoodSecureKey retrieves the value of the environment variable named by "GOOD_SECURE_KEY".
func GoodSecureKey() string {
	return os.Getenv("GOOD_SECURE_KEY")
}