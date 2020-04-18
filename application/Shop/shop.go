package main

// Shop domain
type Shop struct {
	Name string
}

// NewShop returns a Shop struct
func NewShop(name string) *Shop {
	return &Shop{name}
}
