package storage

import "velocity-limits/internal/models"

type Storage struct {
	accounts     map[string]*models.CustomerAccount
	transactions map[string]struct{}
}

// Return a new Storage struct with default values
func NewStorage() *Storage {
	return &Storage{
		accounts:     make(map[string]*models.CustomerAccount),
		transactions: make(map[string]struct{}),
	}
}

func (s *Storage) GetAccount(customerID string) *models.CustomerAccount {
	if acc, ok := s.accounts[customerID]; ok {
		return acc
	}
	return nil
}

func (s *Storage) AddAccount(account *models.CustomerAccount) *models.CustomerAccount {
	s.accounts[account.CustomerID] = account
	return account
}

// Adds transaction info with (Load ID + Customer ID) key for duplicate detection.
func (s *Storage) AddTransaction(id, customerID string) {
	s.transactions[id+customerID] = struct{}{}
}

func (s *Storage) IsDuplicateTransaction(id, customerID string) bool {
	if _, ok := s.transactions[id+customerID]; ok {
		return true
	}
	return false
}
