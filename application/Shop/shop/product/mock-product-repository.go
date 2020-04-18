package product

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/MihaiBlebea/OnlineShop/Shop/env"
)

// MockRepository for testing
type MockRepository struct {
	products []Product
}

// Add inserts a new document in the roducts collection
func (r *MockRepository) Add(product *Product) error {
	product.ID = primitive.NewObjectID()
	r.products = append(r.products, *product)
	return nil
}

// All returns all documents in the collection
func (r *MockRepository) All() ([]Product, error) {
	return r.products, nil
}

// UpdateQuantity updates the quantity of the product in stock
func (r *MockRepository) UpdateQuantity(product *Product) error {
	for _, prod := range r.products {
		if prod.ID.Hex() == product.ID.Hex() {
			prod.Quantity = product.Quantity
		}
	}
	return nil
}

// FindByPriceAndRating finds a Product with a fix price and sorted by rating
func (r *MockRepository) FindByPriceAndRating(price float64) ([]Product, error) {
	result := []Product{}
	sort.Slice(r.products, func(i, j int) bool {
		return r.products[i].Rating < r.products[j].Rating
	})

	for _, prod := range r.products {
		if prod.Price < price {
			result = append(result, prod)
		}
	}
	return result, nil
}

// FindLowStock returns products that have low quantity
func (r *MockRepository) FindLowStock() ([]Product, error) {
	result := []Product{}
	for _, prod := range r.products {
		if prod.Quantity < 5 {
			result = append(result, prod)
		}
	}
	return result, nil
}

func (r *MockRepository) migrate() error {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	prods, err := fetchData(basepath + "/../../products.json")
	if err != nil {
		return err
	}
	supplyLimit, err := strconv.Atoi(env.Get("SHOP_SUPPLY_LIMIT", "5"))
	if err != nil {
		return err
	}
	for _, prod := range prods {
		prod.ID = primitive.NewObjectID()
		prod.Quantity = supplyLimit
		r.products = append(r.products, prod)
	}
	return nil
}

// NewMockRepository constructs and returns a new MockRepository struct
func NewMockRepository() *MockRepository {
	return &MockRepository{}
}

func fetchData(filePath string) ([]Product, error) {
	products := []Product{}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return products, err
	}

	json.Unmarshal([]byte(file), &products)
	return products, nil
}

// NewMigratedMock returns a pre migrated mock repository
func NewMigratedMock() *MockRepository {
	mock := &MockRepository{}
	mock.migrate()
	return mock
}
