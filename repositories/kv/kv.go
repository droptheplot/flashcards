package kv

import "github.com/go-redis/redis"

type Repository struct {
	DB *redis.Client
}
