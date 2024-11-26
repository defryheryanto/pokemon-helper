package main

import (
	"github.com/defry256/pokemon-helper/config"
	"github.com/go-redis/redis/v8"
)

func setupRedis() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Network:  config.RedisNetwork(),
		Addr:     config.RedisAddress(),
		DB:       config.RedisDB(),
		Username: config.RedisUsername(),
		Password: config.RedisPassword(),
	})

	return redisClient
}
