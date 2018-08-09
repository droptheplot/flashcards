package kv

import "github.com/go-redis/redis"

type Repository struct {
	Client *redis.Client
}
