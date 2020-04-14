package main

// Product is the model for parsing the product json
type Product struct {
	ID       string  `json:"id"`
	Title    string  `json:"title"`
	ProdType string  `json:"type"`
	Price    float64 `json:"price"`
	Rating   int     `json:"rating"`
}
