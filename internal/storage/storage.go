// Package storage defines a Storage struct for storing customer accounts
// and transactions to detect duplication.

// Note that for abstraction Storage struct fields start with lowercase letters
// and because of that it cannot be exported.
package storage

import "velocity-limits/internal/models"

// Storage struct represents customer accounts and processed transactions.
// Fields can't be exported out of package.
type Storage struct {
	accounts     map[string]*models.CustomerAccount
	transactions map[string]struct{}
}

// Returns a new Storage struct with default values
func NewStorage() *Storage {
	return &Storage{
		accounts:     make(map[string]*models.CustomerAccount),
		transactions: make(map[string]struct{}),
	}
}

// Returns a customer account by customer ID.
func (s *Storage) GetAccount(customerID string) *models.CustomerAccount {
	if acc, ok := s.accounts[customerID]; ok {
		return acc
	}
	return nil
}

// Add customer account to the storage with its velocity limits.
func (s *Storage) AddAccount(account *models.CustomerAccount) *models.CustomerAccount {
	s.accounts[account.CustomerID] = account
	return account
}

// Adds transaction info with (Load ID + Customer ID) key for duplicate detection.
func (s *Storage) AddTransaction(id, customerID string) {
	s.transactions[id+customerID] = struct{}{}
}

// Checks for duplicate transaction by load ID and customer ID.
func (s *Storage) IsDuplicateTransaction(id, customerID string) bool {
	if _, ok := s.transactions[id+customerID]; ok {
		return true
	}
	return false
}
