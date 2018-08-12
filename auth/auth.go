package auth

import (
	"crypto/sha256"
	"errors"
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func HashPassword(password string) string {
	h := sha256.New()

	h.Write([]byte(password))

	return fmt.Sprintf("%x", h.Sum(nil))
}

func GenerateToken(userID int, secret string) (string, error) {
	jwtToken := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userID": userID,
		},
	)

	return jwtToken.SignedString([]byte(secret))
}

func ParseToken(token string, secret string) (int, error) {
	jwtToken, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", jwtToken.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		return int(claims["userID"].(float64)), nil
	}

	return 0, errors.New("invalid token")
}
