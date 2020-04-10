package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/go-redis/redis"
)

const productlistKey = "products"

// ProductRepository acts like a layer between the domain and the persistence layer
type ProductRepository struct {
	connection *redis.Client
}

// GetList returns a list of all the products saved which represent keys to query by individual product
func (pr *ProductRepository) GetList() ([]string, error) {
	products := pr.connection.SMembers(productlistKey)
	list, err := products.Result()
	if err != nil {
		return []string{}, err
	}
	return list, nil
}

// FindByMoney returns a product with the highest rating that the customer can afford in his budget
func (pr *ProductRepository) FindByMoney(money float64) (Product, error) {
	list, err := pr.GetList()
	if err != nil {
		return Product{}, err
	}

	products := []Product{}
	for _, key := range list {
		prod, err := pr.connection.HGetAll(key).Result()
		if err != nil {
			//
		}
		price, err := strconv.ParseFloat(prod["Price"], 64)
		if err != nil {
			return Product{}, err
		}

		rating, err := strconv.Atoi(prod["Rating"])
		if err != nil {
			return Product{}, err
		}

		quantity, err := strconv.Atoi(prod["Quantity"])
		if err != nil {
			return Product{}, err
		}

		product := Product{prod["ID"], prod["Title"], price, rating, quantity}
		fmt.Println(product)
		products = append(products, product)
	}

	return Product{}, nil
}

func genKey(input string) string {
	return strings.Replace(strings.ToLower(input), " ", ":", -1)
}

func find(list []string, value string) (int, bool) {
	for i, item := range list {
		if item == value {
			return i, true
		}
	}
	return -1, false
}

// NewProductRepo constructs and returns a new ProductRepository struct
func NewProductRepo(connection *redis.Client) *ProductRepository {
	return &ProductRepository{connection}
}
