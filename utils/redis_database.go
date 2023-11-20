package utils

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
)

var RedisClient *redis.Client

func RedisConnection() error {
	errLoad := LoadEnv()
	if errLoad != nil {
		return errLoad
	}
	url := os.Getenv("REDIS_URL")
	port := os.Getenv("REDIS_PORT")
	//redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	address := url + ":" + port
	fmt.Println("?????????????????????????????????????")
	fmt.Println(address)
	RedisClient = redis.NewClient(&redis.Options{
		Addr: address,
	})
	fmt.Println(RedisClient)
	return nil
}
func AddToRedis() error {
	err := RedisClient.Set(context.Background(), "testing", "hi there", 0).Err()
	if err != nil {
		return err
	}
	value, errorget := RedisClient.Get(context.Background(), "testing").Result()
	if errorget != nil {
		return errorget
	} else {
		fmt.Println(value)
	}

	return nil
}
