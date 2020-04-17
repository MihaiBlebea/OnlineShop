package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Customer model
type Customer struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Name  string             `json:"name"`
	Money float64            `json:"money"`
	Spent float64            `json:"spent"`
	Cart  []Product          `json:"cart"`
}

// NewCustomer returns a new Customer model
func NewCustomer() *Customer {

	surnames, err := fetchData("surnames.json")
	if err != nil {
		fmt.Println(err)
	}
	surname := surnames[genRandom(0, len(surnames)-1)]

	firstnames, err := fetchData("firstnames_m.json")
	if err != nil {
		fmt.Println(err)
	}
	firstname := firstnames[genRandom(0, len(firstnames)-1)]
	fullname := fmt.Sprintf("%s %s", firstname, surname)

	max, err := strconv.Atoi(getenv("CUSTOMER_MONEY", "200"))
	if err != nil {
		fmt.Println(err)
	}
	money := genRandomMoney(0, max)
	return &Customer{Name: fullname, Money: money, Cart: []Product{}}
}

// AddProduct adds a Product to the Customer cart
func (c *Customer) AddProduct(product Product) {
	c.Cart = append(c.Cart, product)
}

// TotalSpent set the Spent attribute on the model
func (c *Customer) TotalSpent(amount float64) {
	c.Spent = amount
}

func genRandom(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func genRandomMoney(min, max int) float64 {
	return float64(genRandom(min, max))
}
