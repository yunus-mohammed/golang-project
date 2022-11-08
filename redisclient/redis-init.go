package redisclient

import "github.com/go-redis/redis/v8"

var RedisClient *redis.Client

func ClientInit() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

}

func GetRedisClient() *redis.Client {
	return RedisClient
}
