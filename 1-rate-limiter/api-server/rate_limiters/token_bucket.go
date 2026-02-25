package rate_limiters

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func TokenBucket(client *redis.Client) gin.HandlerFunc {
	return func(c *gin.Context) {
		decrRes, err := client.Decr(c.Request.Context(), "counter").Result()
		if err != nil {
			c.JSON(500, gin.H{
				"message": "It's not you, it's us",
			})
			c.Abort()
			return
		}

		if decrRes < 0 {
			client.Incr(c.Request.Context(), "counter")
			c.JSON(429, gin.H{
				"message": "Too Many Requests",
			})
			c.Abort()
			return
		}
		fmt.Println("Decremented", decrRes)
		c.Next()
	}
}
