package cache

import (
	"sync"
	"time"

	"spotify/helper/common"
)

//go:generate mockgen -destination=./mocks/mock_$GOFILE -source=$GOFILE -package=mocks
type ServerCache interface {
	GetCode(key string) ([]byte, error)
	SetCode(data interface{}, key string, exp time.Duration) error
}

func NewServerCacheRedis(jedisClient RedisClient) ServerCache {
	instance := &redisCacheImpl{
		redisClient: jedisClient,
	}
	kvStoreInstance.Store(common.SERVERCACHE, instance)
	return instance
}

var (
	kvStoreInstance sync.Map
)

func GetServerCacheInstance(name string) ServerCache {
	var inst, ok = kvStoreInstance.Load(name)
	if ok {
		return inst.(ServerCache)
	}
	return nil
}
