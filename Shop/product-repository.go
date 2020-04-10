package main

import (
	"reflect"
	"strings"

	"github.com/go-redis/redis/v7"
	"github.com/google/uuid"
)

const listKey = "products"

// ProductRepository acts like a layer between the domain and the persistence layer
type ProductRepository struct {
	connection *redis.Client
	listKey    string
}

// Add saves one product to the products collection
func (pr *ProductRepository) Add(product *Product) error {
	// generate a key for the product
	key := genKey(product.Title)

	// Check if this product is already stored in redis
	list, err := pr.GetList()
	if err != nil {
		return err
	}

	// If product exits, then increment quantity and exit
	_, exists := find(list, key)
	if exists == true {
		pr.connection.HIncrBy(key, "quantity", 1)
		return nil
	}

	// Add id to the product
	id := uuid.New().String()
	product.AddID(id)

	// Get all model keys and set them in redis hash
	model := reflect.ValueOf(product).Elem()
	for i := 0; i < model.NumField(); i++ {
		name := model.Type().Field(i).Name
		pr.connection.HSet(key, genKey(name), model.Field(i).Interface())
	}

	// Set default quantty equal to 1
	pr.connection.HSet(key, "quantity", 1)

	// Add the new product to the list of products
	pr.connection.SAdd(pr.listKey, key)

	return nil
}

// GetList returns a list of all the products saved which represent keys to query by individual product
func (pr *ProductRepository) GetList() ([]string, error) {
	products := pr.connection.SMembers(listKey)
	list, err := products.Result()
	if err != nil {
		return []string{}, err
	}
	return list, nil
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
	return &ProductRepository{connection, listKey}
}
