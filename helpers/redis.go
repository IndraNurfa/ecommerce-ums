package helpers

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var RedisClient *redis.Client

func SetupRedis() {
	client := redis.NewClient(&redis.Options{
		Addr:     GetEnv("REDIS_HOST", "localhost:6379"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	defer client.Close()

	ping, err := client.Ping(ctx).Result()

	if err != nil {
		Logger.Error("Failed to connect redis: " + ping)
	}
	Logger.Info("PING REDIS: " + ping)
}
