package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	Tokens         = 3
	RefillInterval = 1 * time.Second
)

func startRefiller(client *redis.Client) {
	ticker := time.NewTicker(RefillInterval)
	defer ticker.Stop()

	for range ticker.C {
		err := client.IncrBy(context.Background(), "counter", Tokens).Err()
		if err != nil {
			fmt.Println("Error refilling tokens:", err)
			continue
		}
		fmt.Printf("[REFILL] %s → counter reset to %d\n", time.Now().Format("15:04:05"), Tokens)
	}
}
