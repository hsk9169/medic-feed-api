package Config

import (
	"fmt"
	"os"
	"github.com/go-redis/redis/v7"
)

var CACHE *redis.Client

func CacheOptions() *redis.Options {
	HOST := os.Getenv("CACHE_HOST")
	PORT := os.Getenv("CACHE_PORT")

	return &redis.Options {
		Addr: fmt.Sprintf("%s:%s", HOST, PORT),
		Password: "",
		DB: 0,
	}
}