package redisclient

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func ClientInit() {

	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println("redis not co0nnected")
		panic(err.Error())
	}

}

func GetRedisClient() *redis.Client {
	return redisClient
}

func RedisClientSetVal(key string, str string) {

	var ctx = context.Background()

	err := GetRedisClient().Set(ctx, key, str, 0).Err()
	if err != nil {
		fmt.Println(err)
	}

}

func RedisClientGetVal(key string) (string, error) {

	var ctx = context.Background()

	val, err := GetRedisClient().Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
	}

	return val, err

}
