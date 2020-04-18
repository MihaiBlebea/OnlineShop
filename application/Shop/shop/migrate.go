package shop

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

	"github.com/MihaiBlebea/OnlineShop/Shop/env"
	"github.com/MihaiBlebea/OnlineShop/Shop/shop/product"
)

// MigrateDB runs the migration for the db
func (s *Shop) MigrateDB(filePath string) error {
	prods, err := fetchData(filePath)
	if err != nil {
		return err
	}
	supplyLimit, err := strconv.Atoi(env.Get("SHOP_SUPPLY_LIMIT", "5"))
	if err != nil {
		return err
	}
	productRepo := product.NewRepository()
	for _, prod := range prods {
		prod.Quantity = supplyLimit
		err := productRepo.Add(&prod)
		if err != nil {
			return err
		}
	}
	return nil
}

func fetchData(filePath string) ([]product.Product, error) {
	products := []product.Product{}
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return products, err
	}

	json.Unmarshal([]byte(file), &products)
	return products, nil
}
