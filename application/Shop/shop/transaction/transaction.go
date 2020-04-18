package transaction

import (
	"time"

	"github.com/MihaiBlebea/OnlineShop/Shop/shop/product"
)

const (
	shopName       = "SHOP"
	supplierName   = "SUPPLIER"
	customerName   = "CUSTOMER"
	dateTimeFormat = "2006-01-02 15:04:05"
)

// Transaction domain
type Transaction struct {
	Money     float64
	Buyer     string
	Seller    string
	Products  *[]product.Product
	Timestamp string
}

// IsBuy returns boolean based on the transaction Buyer and Seller
func (t *Transaction) IsBuy() bool {
	if t.Buyer == shopName {
		return true
	}
	return false
}

// NewIn returns a Transaction struct
func NewIn(Money float64, Products *[]product.Product) *Transaction {
	created := time.Now().Format(dateTimeFormat)
	return &Transaction{Money, shopName, supplierName, Products, created}
}

// NewOut returns a Transaction struct
func NewOut(Money float64, Products *[]product.Product) *Transaction {
	created := time.Now().Format(dateTimeFormat)
	return &Transaction{Money, customerName, shopName, Products, created}
}
