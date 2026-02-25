package main

import (
	"fmt"
	"rate_limiter/handlers"
	"rate_limiter/rate_limiters"
	common_redis "rate_limiter/redis"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func main() {
	r := gin.Default()

	redisClient := common_redis.SetupRedis()
	defer func(redisClient *redis.Client) {
		err := redisClient.Close()
		if err != nil {

		}
	}(redisClient)

	r.Use(rate_limiters.TokenBucket(redisClient))

	r.GET("/hello", handlers.HelloHandler)

	if err := r.Run(":5000"); err != nil {
		fmt.Println(err)
	}
}
