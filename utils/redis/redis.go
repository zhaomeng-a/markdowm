package redis

import (
	"fmt"
	"github.com/go-redis/redis/v9"
	"markdown/utils/config"
)

type Pool = redis.Client

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.GC.Redis.Host, config.GC.Redis.Port),
		Password: config.GC.Redis.Password,
		DB:       0,
	})
	return client
}
