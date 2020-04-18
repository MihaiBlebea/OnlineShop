package transaction

import (
	"context"

	"github.com/MihaiBlebea/OnlineShop/Shop/env"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Transactions Collection name for mongo db
const transColName = "transactions"

// Database for mongo db
const dbName = "shop"

// Repository persistence layer for Transaction domain
type Repository struct {
}

// IRepository interface for transaction repository
type IRepository interface {
	Add(transaction *Transaction) error
	All() ([]Transaction, error)
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

// Add saves a Shop struct in the database
func (r Repository) Add(transaction *Transaction) error {
	client, err := r.client()
	if err != nil {
		return err
	}
	ctx := context.TODO()
	defer client.Disconnect(ctx)

	db := client.Database(dbName)
	transactionCollection := db.Collection(transColName)

	_, err = transactionCollection.InsertOne(ctx, transaction)
	if err != nil {
		return err
	}
	return nil
}

// All returns all transactions
func (r Repository) All() ([]Transaction, error) {
	transactions := []Transaction{}

	client, err := r.client()
	if err != nil {
		return transactions, err
	}
	ctx := context.TODO()
	defer client.Disconnect(ctx)

	db := client.Database(dbName)
	transactionCollection := db.Collection(transColName)

	cursor, err := transactionCollection.Find(ctx, bson.D{})
	if err != nil {
		return transactions, err
	}

	err = cursor.All(ctx, &transactions)
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}

// NewRepository returns a new Repository struct
func NewRepository() Repository {
	return Repository{}
}
