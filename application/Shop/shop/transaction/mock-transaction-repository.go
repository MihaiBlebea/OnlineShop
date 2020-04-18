package transaction

// MockRepository for testing
type MockRepository struct {
	transactions []Transaction
}

// Add saves a Shop struct in the database
func (r *MockRepository) Add(transaction *Transaction) error {
	r.transactions = append(r.transactions, *transaction)
	return nil
}

// All returns all transactions
func (r *MockRepository) All() ([]Transaction, error) {
	return r.transactions, nil
}

// NewMockRepository returns a new Repository struct
func NewMockRepository() *MockRepository {
	return &MockRepository{}
}
