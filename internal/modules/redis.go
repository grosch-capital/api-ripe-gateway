package modules

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func redisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "redis.lerigos.com:6379",
		DB:   0,
	})
	return client
}

func GetElementByKey(key string) string {
	client := redisClient()
	defer client.Close()

	ctx := context.Background()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return ""
	}

	return val
}

func SetElementByKey(key string, value string) error {
	client := redisClient()
	defer client.Close()

	ctx := context.Background()

	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}

	return nil
}
