package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClusterClient struct {
	client *redis.ClusterClient
}

func NewRedisClusterClient(client *redis.ClusterClient) RedisClient {
	return &RedisClusterClient{client: client}
}

func (r *RedisClusterClient) HSet(ctx context.Context, key, field string, value interface{}) *redis.IntCmd {
	return r.client.HSet(ctx, key, field, value)
}

func (r *RedisClusterClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	return r.client.HGet(ctx, key, field)
}

func (r *RedisClusterClient) Set(ctx context.Context, key string, value interface{}, exp time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, exp)
}

func (r *RedisClusterClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

func (r *RedisClusterClient) HDel(ctx context.Context, key, field string) *redis.IntCmd {
	return r.client.HDel(ctx, key, field)
}

func (r *RedisClusterClient) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.LPush(ctx, key, values...)
}

func (r *RedisClusterClient) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.RPush(ctx, key, values...)
}

func (r *RedisClusterClient) LPop(ctx context.Context, key string) *redis.StringCmd {
	return r.client.LPop(ctx, key)
}

func (r *RedisClusterClient) Del(ctx context.Context, key string) *redis.IntCmd {
	return r.client.Del(ctx, key)
}
