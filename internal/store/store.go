package store

import (
	"context"
	"time"
)

type Store interface {
	Incr(ctx context.Context, key string, ttl time.Duration) (int64, error)
	Expire(ctx context.Context, key string, ttl time.Duration) error
}
