package database

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client
var ctx = context.Background()

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	if err := redisClient.Ping(ctx).Err(); err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	fmt.Println("Connected to Redis!")
}
