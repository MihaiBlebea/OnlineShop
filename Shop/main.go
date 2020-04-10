package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v7"
)

func main() {
	products, err := fetchData("products.json")
	if err != nil {
		log.Panic(err)
	}
	client := newRedisClient("redis", "6379")

	repo := NewProductRepo(client)

	// Loop
	for {
		index := genRandom(0, len(products))
		product := products[index]
		fmt.Println(product)
		repo.Add(&product)
		time.Sleep(3 * time.Second)
	}
}

func newRedisClient(host, port string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})
	return client
}

func genRandom(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func fetchData(filePath string) ([]Product, error) {
	products := []Product{}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return products, err
	}

	json.Unmarshal([]byte(file), &products)
	return products, nil
}
