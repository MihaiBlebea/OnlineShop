package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// Transactions Collection name for mongo db
const transColName = "transactions"

// Database for mongo db
const dbName = "shop"

// TransactionRepository persistence layer for Transaction domain
type TransactionRepository struct {
}

// ITransactionRepository interface for transaction repository
type ITransactionRepository interface {
	Add(transaction *Transaction) error
	All() ([]Transaction, error)
}

// Add saves a Shop struct in the database
func (tr *TransactionRepository) Add(transaction *Transaction) error {
	client, err := newMongoConnection()
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
func (tr *TransactionRepository) All() ([]Transaction, error) {
	transactions := []Transaction{}

	client, err := newMongoConnection()
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

// NewTransactionRepo returns a new TransactionRepository struct
func NewTransactionRepo() *TransactionRepository {
	return &TransactionRepository{}
}
