package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

func main() {

	// Create customer
	// customer := NewCustomer()
	client := newRedisClient("redis", "6379")
	productRepo := NewProductRepo(client)

	for {
		_, err := productRepo.FindByMoney(200.00)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(10 * time.Second)
	}

	// Check if product quantity > 0

	// Remove product from product repo

	// Add product to customer cart

	// Save customer in customer repo
}

func newRedisClient(host, port string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})
	return client
}
