package redisdb

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func RedisInit() *redis.Client {
	connection := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONNECTION_STRING"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	pong, err := connection.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Redis connection failed", err)
	}
	log.Println("Redis Connected!", "Ping", pong)
	redisClient = connection
	return redisClient
}
