package main

import (
	_ "github.com/EDDYCJY/go-gin-example/pkg/setting"

	"github.com/go-redis/redis"
)

var client *redis.Client

func initRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// Set a key/value
func saveToRedis(key string, token string) error {
	_, err := client.Set(key, token, 0).Result()
	if err != nil {
		return err
	}

	return nil
}

func checkKeyExistanceInRedis(key string) bool {
	var output = client.Exists(key).Val()
	if output == 1 {
		return true
	}
	return false
}
