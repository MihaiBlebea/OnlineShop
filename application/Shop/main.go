package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Populate database from json file
	products, err := fetchData("products.json")
	if err != nil {
		log.Println(err)
	}
	productRepo := NewProductRepo()
	productRepo.AddMany(products)

	// Http server
	serve()
}

func find(list []string, value string) (int, bool) {
	for i, item := range list {
		if item == value {
			return i, true
		}
	}
	return -1, false
}

func roundTwoDecimals(input float64) float64 {
	return math.Round(input*100) / 100
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

func newMongoConnection() (*mongo.Client, error) {
	mongoHost := getenv("MONGO_HOST", "localhost")
	mongoPort := getenv("MONGO_PORT", "27016")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + mongoHost + ":" + mongoPort))
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		fmt.Println(fmt.Sprintf("Could not find env for key %s. Returning default value %s", key, fallback))
		return fallback
	}
	return value
}
