package modules

import (
	"context"
	"crypto/tls"

	"github.com/redis/go-redis/v9"
)

func redisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "redis.grosch.capital:6379",
		DB:   1,
		TLSConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
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
