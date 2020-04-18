package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"time"

	cron "github.com/robfig/cron/v3"
)

func main() {
	c := cron.New()

	// Supply shop as first action
	supplyShop()

	c.AddFunc("*/30 * * * *", supplyShop)
	c.AddFunc("* * * * *", orderCustomer)

	c.Start()

	for {
		fmt.Println("Worker is running")
		time.Sleep(30 * time.Minute)
	}
}

func prettyStartPrint(name string) time.Time {
	now := time.Now()
	fmt.Println(fmt.Sprintf("%s - %s: job has started", now.Format("2006-01-02 15:04:05"), name))
	return now
}

func prettyEndPrint(name string, startTime time.Time) {
	now := time.Now()
	duration := now.Sub(startTime)
	fmt.Println(fmt.Sprintf("%s - %s: job has ended. Duration %s", now.Format("2006-01-02 15:04:05"), name, duration.String()))
}

func supplyShop() {
	jobName := "SUPPLY_SHOP"
	startTime := prettyStartPrint(jobName)

	_, err := http.Post("http://shop:8000/supply", "application/json", bytes.NewBuffer(nil))
	if err != nil {
		log.Println(err)
	}
	prettyEndPrint(jobName, startTime)
}

func orderCustomer() {
	jobName := "ORDER_CUSTOMER"
	startTime := prettyStartPrint(jobName)

	_, err := http.Post("http://customer:8000/customer", "application/json", bytes.NewBuffer(nil))
	if err != nil {
		log.Println(err)
	}

	prettyEndPrint(jobName, startTime)
}
