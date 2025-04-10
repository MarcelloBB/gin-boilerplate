package db

import (
	"context"
	"fmt"
	"time"

	"github.com/MarcelloBB/gin-boilerplate/config"
	"github.com/redis/go-redis/v9"
)

var (
	Ctx         = context.Background()
	RedisClient *redis.Client
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.LoadConfigIni("redis", "host", "localhost:6379").(string),
		Password: config.LoadConfigIni("redis", "password", "").(string),
		DB:       config.LoadConfigIni("redis", "db", 0).(int),
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error connecting to Redis: %v", err))
	}
}

func SetCacheValue(ctx context.Context, key string, value interface{}) error {
	expiration := time.Duration(config.LoadConfigIni("redis", "expiration", 10).(int)) * time.Minute
	err := RedisClient.Set(Ctx, key, value, expiration).Err()
	if err != nil {
		fmt.Println("Error setting value:", err)
		return err
	}
	return nil
}

func GetCacheValue(ctx context.Context, key string) (string, error) {
	value, err := RedisClient.Get(Ctx, key).Result()
	if err != nil {
		fmt.Println("Error getting value:", err)
		return "", err
	}
	return value, nil
}
