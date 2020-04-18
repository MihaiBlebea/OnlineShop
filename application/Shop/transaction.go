package main

import "time"

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
	Products  *[]Product
	Timestamp string
}

// IsBuy returns boolean based on the transaction Buyer and Seller
func (t *Transaction) IsBuy() bool {
	if t.Buyer == shopName {
		return true
	}
	return false
}

// NewTransactionIn returns a Transaction struct
func NewTransactionIn(Money float64, Products *[]Product) *Transaction {
	created := time.Now().Format(dateTimeFormat)
	return &Transaction{Money, shopName, supplierName, Products, created}
}

// NewTransactionOut returns a Transaction struct
func NewTransactionOut(Money float64, Products *[]Product) *Transaction {
	created := time.Now().Format(dateTimeFormat)
	return &Transaction{Money, customerName, shopName, Products, created}
}
