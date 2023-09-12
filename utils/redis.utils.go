package common_utils

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func InitializeConnection(host, password string, db int) error {
	Client = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})

	// Ping the Redis server to check the connection
	ctx := context.Background()
	_, err := Client.Ping(ctx).Result()
	fmt.Println("Redis server connected...")
	if err != nil {
		return fmt.Errorf("error connecting to Redis: %v", err)
	}

	return nil
}

func SetKey(key, value string) error {
	err := Client.Set(context.Background(), key, value, 0).Err()
	return err
}

func GetKey(key string) (string, error) {
	value, err := Client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("Key not found")
	} else if err != nil {
		return "", err
	}
	return value, nil
}

func DeleteKey(key string) error {
	err := Client.Del(context.Background(), key).Err()
	return err
}

func ClearCache() error {
	err := Client.FlushAll(context.Background()).Err()
	return err
}
