package product

import "go.mongodb.org/mongo-driver/bson/primitive"

// Product is the model for parsing the product json
type Product struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Title       string             `json:"title"`
	ProdType    string             `json:"type"`
	Description string             `json:"description"`
	Filename    string             `json:"filename"`
	Price       float64            `json:"price"`
	Rating      int                `json:"rating"`
	Quantity    int                `json:"quantity"`
}

// GetImage returns the image for the product based on te Filename and github data
func (p *Product) GetImage() string {
	return "https://github.com/wedeploy-examples/supermarket-web-example/blob/master/ui/assets/images/" + p.Filename + "?raw=true"
}

// IncrementQuantity adds 1 item to the product quatity
func (p *Product) IncrementQuantity() {
	p.Quantity++
}

// DecrementQuantity substracts 1 item from the product quatity
func (p *Product) DecrementQuantity() {
	if p.Quantity > 0 {
		p.Quantity--
	}
}

// SetQuantity sets the quantity on the product
func (p *Product) SetQuantity(input int) {
	p.Quantity = input
}

// New returns a new Product struct
func New(title, prodType, description, filename string, price float64, rating int, quality int) *Product {
	return &Product{
		Title:       title,
		ProdType:    prodType,
		Description: description,
		Filename:    filename,
		Price:       price,
		Rating:      rating,
		Quantity:    quality,
	}
}

// WithID returns a new Product struct. Must supply ID
func WithID(id primitive.ObjectID, title, prodType, description, filename string, price float64, rating int, quality int) *Product {
	return &Product{
		ID:          id,
		Title:       title,
		ProdType:    prodType,
		Description: description,
		Filename:    filename,
		Price:       price,
		Rating:      rating,
		Quantity:    quality,
	}
}
