package log

import (
	"encoding/json"
	"time"

	"github.com/MihaiBlebea/OnlineShop/Shop/env"
	"github.com/go-redis/redis"
)

const logsKey = "stream"

// Event is the wrapper struct for logs
type Event struct {
	Service   string      `json:"service"`
	Code      string      `json:"code"`
	Body      interface{} `json:"body"`
	Timestamp string      `json:"timestamp"`
}

func newRedisClient(host, port string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: host + ":" + port,
	})
	return client
}

// Log saves a json to the event stream in Redis
func Log(code string, log interface{}) error {
	client := newRedisClient(env.Get("REDIS_HOST", "localhost"), env.Get("REDIS_PORT", "6379"))
	defer client.Close()

	event := Event{"shop", code, log, time.Now().Format("2006-01-02 15:04:05")}
	logByte, err := json.Marshal(event)
	if err != nil {
		return err
	}
	client.SAdd(logsKey, logByte)
	return nil
}
