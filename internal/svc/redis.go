package svc

import (
	"github.com/Songsuh/go_blog/internal/global"
	"github.com/redis/go-redis/v9"
)

func CreateRedis(c *global.RedisConfig) *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:         c.Addr,
		Username:     c.Username,
		Password:     c.Password,
		DB:           c.Database,
		ClientName:   c.ClientName,
		MaxRetries:   c.MaxRetries,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
		PoolSize:     c.PoolSize,
		MinIdleConns: c.MinIdleConns,
		MaxIdleConns: c.MaxIdleConns,
	})

	return rdb
}
