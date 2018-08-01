package auth

import (
	"crypto/sha256"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

func HashPassword(password string) string {
	h := sha256.New()

	h.Write([]byte(password))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenerateToken() (string, error) {
	u, err := uuid.NewV4()

	return u.String(), err
}
