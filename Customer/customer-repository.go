package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// CustomerRepository acts like a layer between the domain and the persistence layer
type CustomerRepository struct {
}

// Add inserts a new document in the roducts collection
func (cr *CustomerRepository) Add(customer *Customer) (interface{}, error) {
	client, err := newMongoConnection()
	if err != nil {
		return nil, err
	}
	ctx := context.TODO()
	defer client.Disconnect(ctx)

	db := client.Database("customers")
	customerCollection := db.Collection("customers")

	result, err := customerCollection.InsertOne(ctx, customer)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

// All returns all the documents in the collection
func (cr *CustomerRepository) All() ([]Customer, error) {
	customers := []Customer{}

	client, err := newMongoConnection()
	if err != nil {
		return customers, err
	}
	ctx := context.TODO()
	defer client.Disconnect(ctx)

	db := client.Database("customers")
	customerCollection := db.Collection("customers")

	cursor, err := customerCollection.Find(ctx, bson.D{})
	if err != nil {
		return customers, err
	}

	err = cursor.All(ctx, &customers)
	if err != nil {
		return customers, err
	}

	return customers, nil
}

// NewCustomerRepo constructs and returns a new CustomerRepository struct
func NewCustomerRepo() *CustomerRepository {
	return &CustomerRepository{}
}
