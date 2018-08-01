package kv

import (
	"fmt"
	"time"
)

const (
	tokenExpiration = 8760 * time.Hour
	tokenKeyFormat  = "users:%d"
)

func (r *Repository) CreateToken(token string, userID int) error {
	k := fmt.Sprintf(tokenKeyFormat, userID)

	return r.DB.Set(k, token, tokenExpiration).Err()
}
