package main

import (
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer model
type Customer struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Money float64            `json:"money"`
	Spent float64            `json:"spent"`
	Cart  []Product          `json:"cart"`
}

// NewCustomer returns a new Customer model
func NewCustomer() *Customer {
	// id := uuid.New().String()
	money := genRandomMoney(0, 200)
	return &Customer{Money: money, Cart: []Product{}}
}

// AddProduct adds a Product to the Customer cart
func (c *Customer) AddProduct(product Product) {
	c.Cart = append(c.Cart, product)
}

// TotalSpent set the Spent attribute on the model
func (c *Customer) TotalSpent(amount float64) {
	c.Spent = amount
}

func genRandomMoney(min, max int) float64 {
	rand.Seed(time.Now().UnixNano())
	return float64(rand.Intn(max-min+1) + min)
}
