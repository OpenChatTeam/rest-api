package lib

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var (
	RedisClient *redis.Client
)

func InitRedisClient() {

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}
