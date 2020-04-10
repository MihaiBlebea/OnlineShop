package main

// Product is the model for parsing the product json
type Product struct {
	ID       string
	Title    string
	Price    float64
	Rating   int
	Quantity int
}
