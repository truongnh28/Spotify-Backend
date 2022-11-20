package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type RedisClient interface {
	HSet(ctx context.Context, key, field string, values interface{}) *redis.IntCmd
	HGet(ctx context.Context, key, field string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expireTime time.Duration) *redis.StatusCmd
	Get(ctx context.Context, key string) *redis.StringCmd
	HDel(ctx context.Context, key, field string) *redis.IntCmd
	LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	LPop(ctx context.Context, key string) *redis.StringCmd
	RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	Del(ctx context.Context, key string) *redis.IntCmd
}
