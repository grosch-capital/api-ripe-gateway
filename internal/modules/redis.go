package modules

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func redisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "master.redis--nrvfy9mhwj5r.addon.code.run:6379", // Replace with your Redis server address
		Password: "ea602ee160f32eb614cb84a1e9117b04",               // Replace with your Redis password
		DB:       0,                                                // Replace with your Redis database number
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
