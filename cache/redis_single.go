package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisSingleClient struct {
	client *redis.Client
}

func NewRedisSingleClient(client *redis.Client) RedisClient {
	return &RedisSingleClient{client: client}
}

func (r *RedisSingleClient) HSet(ctx context.Context, key, field string, value interface{}) *redis.IntCmd {
	return r.client.HSet(ctx, key, field, value)
}

func (r *RedisSingleClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	return r.client.HGet(ctx, key, field)
}

func (r *RedisSingleClient) Set(ctx context.Context, key string, value interface{}, expireTime time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expireTime)
}

func (r *RedisSingleClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *RedisSingleClient) HDel(ctx context.Context, key, field string) *redis.IntCmd {
	return r.client.HDel(ctx, key, field)
}

func (r *RedisSingleClient) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.LPush(ctx, key, values...)
}

func (r *RedisSingleClient) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.RPush(ctx, key, values...)
}

func (r *RedisSingleClient) LPop(ctx context.Context, key string) *redis.StringCmd {
	return r.client.LPop(ctx, key)
}

func (r *RedisSingleClient) Del(ctx context.Context, key string) *redis.IntCmd {
	return r.client.Del(ctx, key)
}
