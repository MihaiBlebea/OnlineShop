package main

import (
	"github.com/MihaiBlebea/OnlineShop/Shop/http"
	"github.com/MihaiBlebea/OnlineShop/Shop/shop"
	s "github.com/MihaiBlebea/OnlineShop/Shop/shop"
)

func main() {
	shop := s.New(
		"Tesco limited",
		shop.ProductRepository(),
		shop.TransactionRepository(),
	)
	shop.MigrateDB("products.json")

	http := http.New(shop)
	http.Serve()
}
