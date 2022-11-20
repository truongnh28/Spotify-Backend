package cache

import (
	"context"
	"time"

	"spotify/helper/common"

	"github.com/go-redis/redis"
)

// redisCacheImpl implements ServerCache
type redisCacheImpl struct {
	redisClient RedisClient
}

func (redisCache *redisCacheImpl) GetCode(key string) ([]byte, error) {
	data, err := redisCache.redisClient.Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return nil, common.ErrEmptyStoredData
	}
	if len(data) > 0 {
		return data, nil
	}
	return nil, err
}

func (redisCache *redisCacheImpl) SetCode(data interface{}, key string, exp time.Duration) error {
	return redisCache.redisClient.Set(context.Background(), key, data, exp*time.Second).Err()
}
