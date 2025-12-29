package security

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomToken buat string acak Base64
func GenerateRandomToken(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}

	return base64.RawURLEncoding.EncodeToString(b), nil
}
