package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
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
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb-shop:27017"))
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
