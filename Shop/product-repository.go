package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ProductRepository acts like a layer between the domain and the persistence layer
type ProductRepository struct {
}

// Add inserts a new document in the roducts collection
func (pr *ProductRepository) Add(product *Product) error {
	client, err := newMongoConnection()
	if err != nil {
		return err
	}
	ctx := context.TODO()
	defer client.Disconnect(ctx)

	db := client.Database("shop")
	productCollection := db.Collection("products")

	_, err = productCollection.InsertOne(ctx, product)
	if err != nil {
		return err
	}
	return nil
}

// AddMany adds more then one document to collection
func (pr *ProductRepository) AddMany(products []Product) {
	for _, product := range products {
		product.IncrementQuantity()
		pr.Add(&product)
	}
}

// All returns all documents in the collection
func (pr *ProductRepository) All() ([]Product, error) {
	products := []Product{}

	client, err := newMongoConnection()
	if err != nil {
		return products, err
	}
	ctx := context.TODO()
	defer client.Disconnect(ctx)

	db := client.Database("shop")
	productCollection := db.Collection("products")

	cursor, err := productCollection.Find(ctx, bson.D{})
	if err != nil {
		return products, err
	}

	err = cursor.All(ctx, &products)
	if err != nil {
		return products, err
	}

	return products, nil
}

// UpdateQuantity updates the quantity of the product in stock
func (pr *ProductRepository) UpdateQuantity(product *Product) error {
	client, err := newMongoConnection()
	if err != nil {
		return err
	}
	ctx := context.TODO()
	defer client.Disconnect(ctx)

	db := client.Database("shop")
	productCollection := db.Collection("products")

	_, err = productCollection.UpdateOne(
		ctx,
		bson.M{"_id": product.ID},
		bson.D{
			{Key: "$set", Value: bson.D{primitive.E{Key: "quantity", Value: product.Quantity}}},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

// FindByPriceAndRating finds a Product with a fix price and sorted by rating
func (pr *ProductRepository) FindByPriceAndRating(price float64) ([]Product, error) {
	products := []Product{}

	client, err := newMongoConnection()
	if err != nil {
		return products, err
	}
	ctx := context.TODO()
	defer client.Disconnect(ctx)

	db := client.Database("shop")
	productCollection := db.Collection("products")

	opts := options.Find()
	opts.SetSort(bson.D{{Key: "rating", Value: -1}})

	sortCursor, err := productCollection.Find(ctx, bson.D{
		{Key: "price", Value: bson.D{{Key: "$lt", Value: price}}},
		{Key: "quantity", Value: bson.D{{Key: "$gt", Value: 0}}},
	}, opts)
	if err != nil {
		return products, err
	}

	err = sortCursor.All(ctx, &products)
	if err != nil {
		return products, err
	}
	return products, nil
}

// NewProductRepo constructs and returns a new ProductRepository struct
func NewProductRepo() *ProductRepository {
	return &ProductRepository{}
}
