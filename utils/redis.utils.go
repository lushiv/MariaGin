package common_utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func InitializeRedisConnection(host, password string, db int) error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})

	// Ping the Redis server to check the connection
	ctx := context.Background()
	_, err := RedisClient.Ping(ctx).Result()

	if err != nil {
		return fmt.Errorf("error connecting to Redis: %v", err)
	}

	return nil
}

func SetKey(key, value string) error {
	err := RedisClient.Set(context.Background(), key, value, 0).Err()
	return err
}

func GetKey(key string) (string, error) {
	value, err := RedisClient.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Key not found")
	} else if err != nil {
		return "", err
	}
	return value, nil
}

func DeleteKey(key string) error {
	err := RedisClient.Del(context.Background(), key).Err()
	return err
}

func ClearCache() error {
	err := RedisClient.FlushAll(context.Background()).Err()
	return err
}
