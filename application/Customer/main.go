package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Http server
	serve()
}

func newMongoConnection() (*mongo.Client, error) {
	mongoHost := getenv("MONGO_HOST", "localhost")
	mongoPort := getenv("MONGO_PORT", "27017")

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

func fetchData(filePath string) ([]string, error) {
	names := []string{}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return names, err
	}

	json.Unmarshal([]byte(file), &names)
	return names, nil
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		fmt.Println(fmt.Sprintf("Could not find env for key %s. Returning default value %s", key, fallback))
		return fallback
	}
	return value
}
