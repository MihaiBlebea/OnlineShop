package main

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

func main() {
	// go shopLoop()
	// go customerLoop()
	for {
		for i := 0; i < 5; i++ {
			_, err := http.Post("http://shop:8000/supply", "application/json", bytes.NewBuffer(nil))
			if err != nil {
				log.Panic(err)
			}
		}
		time.Sleep(10 * time.Second)

		_, err := http.Post("http://customer:8000/customer", "application/json", bytes.NewBuffer(nil))
		if err != nil {
			log.Panic(err)
		}

		time.Sleep(10 * time.Second)
	}
}

// func shopLoop() {
// 	for {
// 		_, err := http.Post("http://shop:8000/supply", "application/json", bytes.NewBuffer([]byte{}))
// 		if err != nil {
// 			log.Panic(err)
// 		}

// 		time.Sleep(10 * time.Second)
// 	}
// }

// func customerLoop() {
// 	for {
// 		_, err := http.Post("http://customer:8000/customer", "application/json", bytes.NewBuffer([]byte{}))
// 		if err != nil {
// 			log.Panic(err)
// 		}

// 		time.Sleep(10 * time.Second)
// 	}
// }
