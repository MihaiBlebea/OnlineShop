package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const port = ":8000"

func serve() {
	router := httprouter.New()

	router.GET("/", indexHandler)
	router.GET("/products", getProductsHandler)
	router.GET("/account/transactions", transactionHandler)
	router.GET("/account/balance", balanceHandler)
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

	supplied := []Product{}
	productRepo := NewProductRepo()
	transactionRepo := NewTransactionRepo()

	// Low stock products
	lsProducts, err := productRepo.All()
	if err != nil {
		log.Panic(err)
	}

	totalCost := 0.00
	supplyLimit, err := strconv.Atoi(getenv("SHOP_SUPPLY_LIMIT", "5"))
	if err != nil {
		log.Panic(err)
	}

	for _, product := range lsProducts {
		diffQuantity := supplyLimit - product.Quantity
		if diffQuantity > 0 {
			product.SetQuantity(supplyLimit)
			cost := float64(diffQuantity) * product.Price
			totalCost += cost

			suppliedProduct := Product{
				product.ID,
				product.Title,
				product.ProdType,
				product.Description,
				product.Filename,
				product.Price,
				product.Rating,
				diffQuantity,
			}
			supplied = append(supplied, suppliedProduct)

			productRepo.UpdateQuantity(&product)
		}
	}

	// New transaction
	transaction := NewTransactionIn(totalCost, &supplied)
	transactionRepo.Add(transaction)

	err = json.NewEncoder(w).Encode(supplied)
	if err != nil {
		log.Panic(err)
	}

	// Logging
	Log("SHOP_SUPPLIED", supplied)
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
	transactionRepo := NewTransactionRepo()

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

		boughtProduct := Product{
			product.ID,
			product.Title,
			product.ProdType,
			product.Description,
			product.Filename,
			product.Price,
			product.Rating,
			1,
		}

		cart = append(cart, boughtProduct)
		total += product.Price

		product.DecrementQuantity()
		productRepo.UpdateQuantity(&product)
	}

	if total > 0 {
		transaction := NewTransactionOut(total, &cart)
		transactionRepo.Add(transaction)
	}

	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		log.Panic(err)
	}

	// Logging
	Log("SHOP_SOLD", cart)
}

func transactionHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	transactionRepo := NewTransactionRepo()
	transactions, err := transactionRepo.All()
	if err != nil {
		log.Panic(err)
	}

	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		log.Panic(err)
	}
}

func balanceHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	transactionRepo := NewTransactionRepo()
	transactions, err := transactionRepo.All()
	if err != nil {
		log.Panic(err)
	}

	balance := 0.00
	for _, transaction := range transactions {
		if transaction.IsBuy() {
			balance -= transaction.Money
		} else {
			balance += transaction.Money
		}
	}

	err = json.NewEncoder(w).Encode(roundTwoDecimals(balance))
	if err != nil {
		log.Panic(err)
	}
}
