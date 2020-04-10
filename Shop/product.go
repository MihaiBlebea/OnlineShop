package main

// Product is the model for parsing the product json
type Product struct {
	Title       string
	ProdType    string
	Description string
	Filename    string
	Height      int
	Width       int
	Price       float64
	Rating      int
	ID          string
}

// GetImage returns the image for the product based on te Filename and github data
func (p *Product) GetImage() string {
	return "https://github.com/wedeploy-examples/supermarket-web-example/blob/master/ui/assets/images/" + p.Filename + "?raw=true"
}

// AddID adds an UUID to the product struct
func (p *Product) AddID(id string) {
	p.ID = id
}
