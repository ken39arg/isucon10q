package main

import (
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	rdb    *redis.Client
	reonce sync.Once
)

func initRedis() {
	reonce.Do(func() {
		rdb = redis.NewClient(&redis.Options{
			// 1台の場合
			// Network: "unix",
			// Addr: "/var/run/redis/redis-server.sock",
			Addr:     "redis:6379",
			Password: "",
			// MaxRetries: 3, // default no retry
			// PoolSize: 10, // default 10
		})
	})
}
