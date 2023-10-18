package utils

import (
	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

var (
	MyCache *cache.Cache
	host    = "localhost"
)

func init() {
	client := redis.NewClient(&redis.Options{
		Addr: host + ":6379",
	})

	MyCache = cache.New(&cache.Options{
		Redis: client,
	})
}
