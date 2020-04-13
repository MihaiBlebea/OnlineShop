package main

import (
	"bytes"
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const port = ":8000"

func serve() {
	router := httprouter.New()

	router.GET("/", indexHandler)
	router.POST("/customer", createCustomerHandler)
	router.GET("/customers", getCustomersHandler)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Panic(err)
	}
}

func setupHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).WriteHeader(http.StatusOK)
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	health := make(map[string]string)
	health["status"] = "OK"

	jsonHealth, err := json.Marshal(health)
	if err != nil {
		log.Panic(err)
	}
	err = json.NewEncoder(w).Encode(jsonHealth)
	if err != nil {
		log.Panic(err)
	}
}

func createCustomerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	customer := NewCustomer()
	customerRepo := NewCustomerRepo()

	body := map[string]interface{}{"money": customer.Money}

	bodyBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := http.Post("http://shop:8000/order", "application/json", bytes.NewBuffer(bodyBytes))
	if err != nil {
		log.Fatal(err)
	}
	products := []Product{}
	err = json.NewDecoder(response.Body).Decode(&products)
	if err != nil {
		log.Fatal(err)
	}

	total := 0.00
	for _, product := range products {
		customer.AddProduct(product)
		total += product.Price
	}
	total = math.Round(total*100) / 100

	customer.TotalSpent(total)
	id, err := customerRepo.Add(customer)
	if err != nil {
		log.Fatal(err)
	}
	customer.ID = id.(primitive.ObjectID)

	err = json.NewEncoder(w).Encode(customer)
	if err != nil {
		log.Panic(err)
	}

	// Logging
	event := make(map[string]interface{})
	payload := make(map[string]interface{})
	cart := []string{}
	for _, product := range products {
		cart = append(cart, product.ID)
	}

	payload["customer_id"] = customer.ID
	payload["spent"] = total
	payload["money"] = customer.Money
	payload["cart"] = cart
	event["service"] = "customer"
	event["code"] = "CUSTOMER_BOUGHT"
	event["body"] = payload
	Log(event)
}

func getCustomersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	customerRepo := NewCustomerRepo()
	customers, err := customerRepo.All()
	if err != nil {
		log.Panic(err)
	}

	err = json.NewEncoder(w).Encode(customers)
	if err != nil {
		log.Panic(err)
	}
}
