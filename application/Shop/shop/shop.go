package shop

import (
	"strconv"

	"github.com/MihaiBlebea/OnlineShop/Shop/env"
	"github.com/MihaiBlebea/OnlineShop/Shop/shop/product"
	"github.com/MihaiBlebea/OnlineShop/Shop/shop/transaction"
)

// Shop domain
type Shop struct {
	name            string
	productRepo     product.IRepository
	transactionRepo transaction.IRepository
}

// New returns a Shop struct
func New(name string, productRepo product.IRepository, transactionRepo transaction.IRepository) *Shop {
	return &Shop{name, productRepo, transactionRepo}
}

// Supply is called to supply the shop
func (s *Shop) Supply() ([]product.Product, error) {
	supplied := []product.Product{}

	// Low stock products
	lsProducts, err := s.productRepo.All()
	if err != nil {
		return supplied, err
	}

	totalCost := 0.00
	supplyLimit, err := strconv.Atoi(env.Get("SHOP_SUPPLY_LIMIT", "5"))
	if err != nil {
		return supplied, err
	}

	for _, prod := range lsProducts {
		diffQuantity := supplyLimit - prod.Quantity
		if diffQuantity > 0 {
			prod.SetQuantity(supplyLimit)
			cost := float64(diffQuantity) * prod.Price
			totalCost += cost

			suppliedProduct := product.WithID(
				prod.ID,
				prod.Title,
				prod.ProdType,
				prod.Description,
				prod.Filename,
				prod.Price,
				prod.Rating,
				diffQuantity,
			)
			supplied = append(supplied, *suppliedProduct)

			s.productRepo.UpdateQuantity(&prod)
		}
	}

	// New transaction
	transaction := transaction.NewIn(totalCost, &supplied)
	s.transactionRepo.Add(transaction)

	return supplied, nil
}

// Order - customer orders from the shop
func (s *Shop) Order(money float64) ([]product.Product, error) {
	cart := []product.Product{}

	products, err := s.productRepo.FindByPriceAndRating(money)
	if err != nil {
		return cart, err
	}

	total := 0.00
	for _, prod := range products {
		if total+prod.Price >= money {
			break
		}

		boughtProduct := product.WithID(
			prod.ID,
			prod.Title,
			prod.ProdType,
			prod.Description,
			prod.Filename,
			prod.Price,
			prod.Rating,
			1,
		)

		cart = append(cart, *boughtProduct)
		total += prod.Price

		prod.DecrementQuantity()
		s.productRepo.UpdateQuantity(&prod)
	}

	if total > 0 {
		transaction := transaction.NewOut(total, &cart)
		s.transactionRepo.Add(transaction)
	}

	return cart, nil
}

// Transactions returns all transactions
func (s *Shop) Transactions() ([]transaction.Transaction, error) {
	transactions, err := s.transactionRepo.All()
	if err != nil {
		return []transaction.Transaction{}, err
	}
	return transactions, nil
}

// Balance returns a float withthe balance account of the shop
func (s *Shop) Balance() (float64, error) {
	transactions, err := s.transactionRepo.All()
	if err != nil {
		return 0.00, err
	}

	balance := 0.00
	for _, transaction := range transactions {
		if transaction.IsBuy() {
			balance -= transaction.Money
		} else {
			balance += transaction.Money
		}
	}
	return balance, nil
}

// StockProducts returns all products with quantity greater then 1
func (s *Shop) StockProducts() ([]product.Product, error) {
	stockProducts := []product.Product{}

	products, err := s.productRepo.All()
	if err != nil {
		return stockProducts, err
	}

	for _, prod := range products {
		if prod.Quantity > 0 {
			stockProducts = append(stockProducts, prod)
		}
	}
	return stockProducts, nil
}
