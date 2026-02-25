package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	requestCount = 5
)

func sendRequests() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		var wg sync.WaitGroup
		results := make(chan string, requestCount)

		for i := 0; i < requestCount; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()

				resp, err := http.Get("http://localhost:5000/hello")
				if err != nil {
					results <- fmt.Sprintf("  [Request %d] Error: %v", id, err)
					return
				}
				defer resp.Body.Close()

				results <- fmt.Sprintf("  [Request %d] %d %s", id, resp.StatusCode, http.StatusText(resp.StatusCode))
			}(i + 1)
		}

		go func() {
			wg.Wait()
			close(results)
		}()

		fmt.Printf("\n[TICK] %s — sending %d requests\n", time.Now().Format("15:04:05"), requestCount)
		for result := range results {
			fmt.Println(result)
		}
	}
}

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
		Protocol: 2,
	})
	defer client.Close()

	client.Set(context.Background(), "counter", 0, 0).Err()

	fmt.Println("Starting refiller and request sender...")

	go startRefiller(client)
	sendRequests()
}
