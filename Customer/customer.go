package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Customer model
type Customer struct {
	ID    string
	Money float64
	Cart  []Product
}

// NewCustomer returns a new Customer model
func NewCustomer() *Customer {
	id := uuid.New().String()
	money := genRandomMoney(0, 200)
	return &Customer{id, money, []Product{}}
}

// AddProduct adds a Product to the Customer cart
func (c *Customer) AddProduct(product Product) {
	c.Cart = append(c.Cart, product)
}

func genRandomMoney(min, max int) float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(rand.Intn(max-min+1) + min)
}
