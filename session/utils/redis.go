package utils

import (
	"fmt"
	"os"

	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

var (
	MyCache *cache.Cache
)

func init() {
	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	MyCache = cache.New(&cache.Options{
		Redis: client,
	})
}
