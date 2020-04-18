package shop

import (
	"strconv"
	"testing"

	"github.com/MihaiBlebea/OnlineShop/Shop/env"
	"github.com/MihaiBlebea/OnlineShop/Shop/shop/product"
	"github.com/MihaiBlebea/OnlineShop/Shop/shop/transaction"
)

type MickProductRepository struct {
}

// TestShopName the shop name
func TestShopName(t *testing.T) {
	name := "Tesco limited"
	shop := New(
		name,
		ProductRepository(),
		TransactionRepository(),
	)
	if shop.name != name {
		t.Error(
			"expected", name,
			"got", shop.name,
		)
	}
}

func TestShopInvalidSupply(t *testing.T) {
	name := "Tesco limited"
	shop := New(
		name,
		product.NewMigratedMock(),
		transaction.NewMockRepository(),
	)

	products, err := shop.Supply()
	if err != nil {
		t.Error(err)
	}

	if len(products) != 0 {
		t.Error(
			"expected", []product.Product{},
			"got", products,
		)
	}
}

func TestShopValidSupply(t *testing.T) {
	name := "Tesco limited"
	shop := New(
		name,
		product.NewMockRepository(),
		transaction.NewMockRepository(),
	)

	products, err := shop.Supply()
	if err != nil {
		t.Error(err)
	}

	limit, err := strconv.Atoi(env.Get("SHOP_SUPPLY_LIMIT", "5"))
	if err != nil {
		t.Error(err)
	}

	for _, prod := range products {
		if prod.Quantity != limit {
			t.Error(
				"expected", limit,
				"got", prod.Quantity,
			)
		}
	}
}

func TestShopInvalidOrder(t *testing.T) {
	name := "Tesco limited"
	shop := New(
		name,
		product.NewMigratedMock(),
		transaction.NewMockRepository(),
	)
	budget := 0.00

	cart, err := shop.Order(budget)
	if err != nil {
		t.Error(err)
	}

	if len(cart) > 0 {
		t.Error(
			"expected", "Cart length is 0",
			"got", "Cart length is greater then 0",
		)
	}
}

func TestShopValidOrder(t *testing.T) {
	name := "Tesco limited"
	shop := New(
		name,
		product.NewMigratedMock(),
		transaction.NewMockRepository(),
	)
	budget := 200.00

	cart, err := shop.Order(budget)
	if err != nil {
		t.Error(err)
	}

	if len(cart) == 0 {
		t.Error(
			"expected", "Cart length is greater then 0",
			"got", "Cart length is 0",
		)
	}

	total := 0.00
	for _, prod := range cart {
		total += prod.Price
	}

	if total > budget {
		t.Error(
			"expected", "Cart price is lower then the budget",
			"got", "Cart price is equal of greater then the budget",
		)
	}
}

func TestShopStockProducts(t *testing.T) {
	name := "Tesco limited"
	shop := New(
		name,
		product.NewMigratedMock(),
		transaction.NewMockRepository(),
	)

	prods, err := shop.StockProducts()
	if err != nil {
		t.Error(err)
	}

	if len(prods) == 0 {
		t.Error(
			"expected", "Products length is greater then 0",
			"got", "Products length is 0",
		)
	}

	for _, prod := range prods {
		if prod.Quantity == 0 {
			t.Error(
				"expected", "Product should have Quantity greater then 0",
				"got", "Product quantity is 0",
			)
		}
	}
}

func TestShopOrderTransactions(t *testing.T) {
	name := "Tesco limited"
	shop := New(
		name,
		product.NewMigratedMock(),
		transaction.NewMockRepository(),
	)

	transactions, err := shop.Transactions()
	if err != nil {
		t.Error(err)
	}
	if len(transactions) > 0 {
		t.Error(
			"expected", "Transactions length should be 0",
			"got", "Transactions length is greater then 0",
		)
	}

	budget := 200.00
	_, err = shop.Order(budget)
	if err != nil {
		t.Error(err)
	}

	transactions, err = shop.Transactions()
	if err != nil {
		t.Error(err)
	}
	if len(transactions) != 1 {
		t.Error(
			"expected", "Transactions length should be 1",
			"got", "Transactions length is not 1",
		)
	}
	if transactions[0].Buyer != "CUSTOMER" {
		t.Error(
			"expected", "Buyer is CUSTOMER",
			"got", transactions[0].Buyer,
		)
	}
	if transactions[0].Seller != "SHOP" {
		t.Error(
			"expected", "Seller is SHOP",
			"got", transactions[0].Seller,
		)
	}

	balance, err := shop.Balance()
	if err != nil {
		t.Error(err)
	}

	if balance <= 0 {
		t.Error(
			"expected", "Balance is positive",
			"got", balance,
		)
	}
}

func TestShopSupplyTransactions(t *testing.T) {
	name := "Tesco limited"
	shop := New(
		name,
		product.NewMockRepository(),
		transaction.NewMockRepository(),
	)

	transactions, err := shop.Transactions()
	if err != nil {
		t.Error(err)
	}
	if len(transactions) > 0 {
		t.Error(
			"expected", "Transactions length should be 0",
			"got", "Transactions length is greater then 0",
		)
	}

	_, err = shop.Supply()
	if err != nil {
		t.Error(err)
	}

	transactions, err = shop.Transactions()
	if err != nil {
		t.Error(err)
	}
	if len(transactions) != 1 {
		t.Error(
			"expected", "Transactions length should be 1",
			"got", "Transactions length is not 1",
		)
	}
	if transactions[0].Buyer != "SHOP" {
		t.Error(
			"expected", "Buyer is SHOP",
			"got", transactions[0].Buyer,
		)
	}
	if transactions[0].Seller != "SUPPLIER" {
		t.Error(
			"expected", "Seller is SHOP",
			"got", transactions[0].Seller,
		)
	}

	balance, err := shop.Balance()
	if err != nil {
		t.Error(err)
	}

	if balance > 0 {
		t.Error(
			"expected", "Balance is negative",
			"got", balance,
		)
	}
}
