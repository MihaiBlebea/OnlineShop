package main

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis"
)

const redisHost = "redis"
const redisPort = "6379"
const logsKey = "stream"

func newRedisClient(host, port string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})
	return client
}

// Log saves a json to the event stream in Redis
func Log(log map[string]interface{}) error {
	log["timestamp"] = time.Now()
	client := newRedisClient(redisHost, redisPort)
	defer client.Close()

	logByte, err := json.Marshal(log)
	if err != nil {
		return err
	}
	client.SAdd(logsKey, logByte)
	return nil
}
