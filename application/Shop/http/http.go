package http

import (
	"encoding/json"
	"log"
	"net/http"

	event "github.com/MihaiBlebea/OnlineShop/Shop/log"
	"github.com/MihaiBlebea/OnlineShop/Shop/shop"
	"github.com/julienschmidt/httprouter"
)

const port = ":8000"

// HTTP serves the shop over http
type HTTP struct {
	shop *shop.Shop
}

// Serve initializes the http server
func (h *HTTP) Serve() {
	router := httprouter.New()

	router.GET("/", h.indexHandler)
	router.GET("/products", h.getProductsHandler)
	router.GET("/account/transactions", h.transactionHandler)
	router.GET("/account/balance", h.balanceHandler)
	router.POST("/supply", h.supplyHandler)
	router.POST("/order", h.orderHandler)

	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Panic(err)
	}
}

func (h *HTTP) setupHeaders(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
	(*w).WriteHeader(http.StatusOK)
}

func (h *HTTP) indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	health := make(map[string]string)
	health["status"] = "OK"

	err := json.NewEncoder(w).Encode(health)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (h *HTTP) getProductsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	stockProducts, err := h.shop.StockProducts()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = json.NewEncoder(w).Encode(stockProducts)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (h *HTTP) supplyHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	h.setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	supplied, err := h.shop.Supply()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = json.NewEncoder(w).Encode(supplied)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	// Logging
	event.Log("SHOP_SUPPLIED", supplied)
}

func (h *HTTP) orderHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.setupHeaders(&w)
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
		http.Error(w, err.Error(), 500)
	}

	cart, err := h.shop.Order(post.Money)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = json.NewEncoder(w).Encode(cart)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	// Logging
	event.Log("SHOP_SOLD", cart)
}

func (h *HTTP) transactionHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	transactions, err := h.shop.Transactions()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func (h *HTTP) balanceHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.setupHeaders(&w)
	if (*r).Method == "OPTIONS" {
		return
	}

	balance, err := h.shop.Balance()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}

	err = json.NewEncoder(w).Encode(roundTwoDecimals(balance))
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
}

// New builds a new HTTP server
func New(shop *shop.Shop) *HTTP {
	return &HTTP{shop}
}
