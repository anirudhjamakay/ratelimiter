package algorithms

import (
	"context"
	"fmt"
	"time"

	"github.com/anirudhjamakay/ratelimiter/internal/store"
)

type FixedWindow struct {
	store  store.Store
	limit  int64
	window time.Duration
}

func NewFixedWindow(s store.Store, limit int64, window time.Duration) *FixedWindow {
	return &FixedWindow{
		store:  s,
		limit:  limit,
		window: window,
	}
}

func (f *FixedWindow) Allow(ctx context.Context, key string) (bool, error) {
	redisKey := fmt.Sprintf("rate:%s", key)

	count, err := f.store.Incr(ctx, redisKey, f.window)
	if err != nil {
		return false, err
	}

	if count == 1 {
		err = f.store.Expire(ctx, key, f.window)
		if err != nil {
			return false, err
		}
	}

	return count <= f.limit, nil
}
