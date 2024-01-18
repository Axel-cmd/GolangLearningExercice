package db

import "github.com/go-redis/redis/v8"

func NewDatabaseClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis-16939.c274.us-east-1-3.ec2.cloud.redislabs.com:16939",
		Password: "yTKoQvbkBggT0pWoKRbANttRIIeMRS1F",
		DB:       0,
	})
}
