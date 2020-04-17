package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const port = ":8000"

func serve() {
	router := httprouter.New()

	router.GET("/", indexHandler)
	router.GET("/products", getProductsHandler)
	router.POST("/supply", supplyHandler)
	router.POST("/order", orderHandler)

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

func getProductsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	productRepo := NewProductRepo()
	products, err := productRepo.All()
	if err != nil {
		log.Panic(err)
	}

	var inStockProducts []Product
	for _, product := range products {
		if product.Quantity > 0 {
			inStockProducts = append(inStockProducts, product)
		}
	}

	err = json.NewEncoder(w).Encode(inStockProducts)
	if err != nil {
		log.Panic(err)
	}
}

func supplyHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	supplied := []SuppliedProduct{}
	productRepo := NewProductRepo()

	// Low stock products
	lsProducts, err := productRepo.All()
	if err != nil {
		log.Panic(err)
	}

	for _, product := range lsProducts {
		diffQuantity := 5 - product.Quantity
		if diffQuantity > 0 {
			product.SetQuantity(5)
			cost := float64(diffQuantity) * product.Price
			supplied = append(supplied, SuppliedProduct{
				product.ID,
				product.Title,
				product.Quantity,
				diffQuantity,
				roundTwoDecimals(cost),
			})
		}
	}

	err = json.NewEncoder(w).Encode(supplied)
	if err != nil {
		log.Panic(err)
	}

	// Logging
	event := make(map[string]interface{})
	event["service"] = "shop"
	event["code"] = "SHOP_SUPPLIED"
	event["body"] = supplied
	Log(event)
}

func orderHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	type Post struct {
		Money float64 `json:"money"`
	}

	var post Post
	err := decoder.Decode(&post)
	if err != nil {
		panic(err)
	}
	money := post.Money

	productRepo := NewProductRepo()
	products, err := productRepo.FindByPriceAndRating(money)
	if err != nil {
		log.Panic(err)
	}

	total := 0.00
	cart := []Product{}
	for _, product := range products {
		if total+product.Price >= money {
			break
		}
		cart = append(cart, product)
		total += product.Price

		productClone := product
		productClone.DecrementQuantity()
		productRepo.UpdateQuantity(&productClone)
	}

	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		log.Panic(err)
	}

	// Logging
	event := make(map[string]interface{})
	payload := make(map[string]interface{})
	prods := []string{}
	for _, product := range cart {
		prods = append(prods, product.ID.Hex())
	}

	payload["spent"] = math.Round(total*100) / 100
	payload["money"] = money
	payload["cart"] = prods
	event["service"] = "shop"
	event["code"] = "SHOP_SOLD"
	event["body"] = payload
	Log(event)
}
