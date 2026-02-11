package store

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct {
	client *redis.Client
}

func NewRedisStore(addr string) *RedisStore {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		DB:       0,
		Protocol: 2, // important for Redis 7 compatibility
	})

	ctx := context.Background()
	if err := rdb.Ping(ctx).Err(); err != nil {
		panic(fmt.Sprintf("failed to connect to redis: %v", err))
	}

	return &RedisStore{client: rdb}
}

func (r *RedisStore) Expire(ctx context.Context, key string, ttl time.Duration) error {
	err := r.client.Expire(ctx, key, ttl)
	return err.Err()
}

func (r *RedisStore) Incr(ctx context.Context, key string, ttl time.Duration) (int64, error) {
	val, err := r.client.Incr(ctx, key).Result()
	fmt.Println(val)
	if err != nil {
		return 0, err
	}

	return val, nil
}
