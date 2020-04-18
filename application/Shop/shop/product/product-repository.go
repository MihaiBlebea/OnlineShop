package product

import (
	"context"

	"github.com/MihaiBlebea/OnlineShop/Shop/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Repository acts like a layer between the domain and the persistence layer
type Repository struct {
}

// IRepository interface for Repository
type IRepository interface {
	Add(product *Product) error
	All() ([]Product, error)
	UpdateQuantity(product *Product) error
	FindByPriceAndRating(price float64) ([]Product, error)
	FindLowStock() ([]Product, error)
}

func (r Repository) client() (*mongo.Client, error) {
	mongoHost := env.Get("MONGO_HOST", "localhost")
	mongoPort := env.Get("MONGO_PORT", "27016")
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + mongoHost + ":" + mongoPort))
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Add inserts a new document in the roducts collection
func (r Repository) Add(product *Product) error {
	client, err := r.client()
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

// All returns all documents in the collection
func (r Repository) All() ([]Product, error) {
	products := []Product{}

	client, err := r.client()
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
func (r Repository) UpdateQuantity(product *Product) error {
	client, err := r.client()
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
func (r Repository) FindByPriceAndRating(price float64) ([]Product, error) {
	products := []Product{}

	client, err := r.client()
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

// FindLowStock returns products that have low quantity
func (r Repository) FindLowStock() ([]Product, error) {
	products := []Product{}

	client, err := r.client()
	if err != nil {
		return products, err
	}
	ctx := context.TODO()
	defer client.Disconnect(ctx)

	db := client.Database("shop")
	productCollection := db.Collection("products")

	cursor, err := productCollection.Find(ctx, bson.D{{Key: "quantity", Value: bson.D{{Key: "$lt", Value: 5}}}})
	if err != nil {
		return products, err
	}

	err = cursor.All(ctx, &products)
	if err != nil {
		return products, err
	}

	return products, nil
}

// NewRepository constructs and returns a new ProductRepository struct
func NewRepository() Repository {
	return Repository{}
}
