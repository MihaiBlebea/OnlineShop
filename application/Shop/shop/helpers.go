package shop

import (
	"github.com/MihaiBlebea/OnlineShop/Shop/shop/product"
	"github.com/MihaiBlebea/OnlineShop/Shop/shop/transaction"
)

// ProductRepository returns a new product repository struct
func ProductRepository() product.Repository {
	return product.NewRepository()
}

// TransactionRepository returns a new transaction repository struct
func TransactionRepository() transaction.Repository {
	return transaction.NewRepository()
}
