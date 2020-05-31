package session

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"time"
)

var(
	client *redis.Client
)

func New(adds string)  {
	client = redis.NewClient(&redis.Options{
		Addr:adds,
		DB:1,
		DialTimeout:10 * time.Second,
		ReadTimeout:30 * time.Second,
		WriteTimeout:30 * time.Second,
		PoolSize:6,
		PoolTimeout:30 * time.Second,
		MaxRetries:2,
		IdleTimeout:5 * time.Minute,
	})

	pong, err := client.Ping().Result()
	fmt.Print(pong, err)
}

func Set(key, value string, duration time.Duration) (error) {
	err := client.SetNX(key, value, duration).Err()
	return err
}

func Get(key string) (string, error) {
	value, err := client.Get(key).Result()
	return value, err
}