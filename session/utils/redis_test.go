package utils

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/cache/v9"
)

func TestGetCache(t *testing.T) {
	type Object struct {
		Str string
		Num int
	}

	ctx := context.TODO()
	key := "mykey"
	obj := &Object{
		Str: "mystring",
		Num: 42,
	}

	if err := MyCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   key,
		Value: obj,
		TTL:   time.Hour,
	}); err != nil {
		panic(err)
	}

	var wanted Object
	if err := MyCache.Get(ctx, key, &wanted); err == nil {
		fmt.Println(wanted)
	}

	if err := MyCache.Delete(ctx, key); err != nil {
		panic(err)
	}
}
