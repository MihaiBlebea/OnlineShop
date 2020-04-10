package main

import (
	"github.com/go-redis/redis"
)

const listKey = "customers"

// CustomerRepository struct
type CustomerRepository struct {
	connection *redis.Client
}

// Add adds a new Customer to the CustomerRepository
func (cr *CustomerRepository) Add(customer *Customer) error {
	key := customer.ID

	cr.connection.HSet(key, "id", customer.ID)

	// Calculate cart total
	total := 0.0
	cart := ""
	for _, product := range customer.Cart {
		total += product.Price
		cart += product.ID + "|"
	}
	cr.connection.HSet(key, "total", total)
	cr.connection.HSet(key, "cart", cart)

	// Add the new customer to the list of customers
	cr.connection.SAdd(listKey, key)

	return nil
}

// NewCustomerRepo constructs and returns a new CustomerRepository struct
func NewCustomerRepo(connection *redis.Client) *CustomerRepository {
	return &CustomerRepository{connection}
}
