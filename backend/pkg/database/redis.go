package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"os"
)

// NewRedisClient returns a Redis client instance.
func NewRedisClient() *redis.Client {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	return redis.NewClient(&redis.Options{
		Addr: addr,
	})
}

// WithRedisTx runs a Redis transaction (pipeline) with context.
func WithRedisTx(ctx context.Context, client *redis.Client, fn func(pipe redis.Pipeliner) error) ([]redis.Cmder, error) {
	return client.TxPipelined(ctx, fn)
}
