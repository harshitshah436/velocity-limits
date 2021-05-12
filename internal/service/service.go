// Package service implements functions to read input file, process transactions
// from it and write responses to the output file.
package service

import (
	"bufio"
	"encoding/json"
	"log"
	"velocity-limits/config"
	"velocity-limits/internal/models"
	"velocity-limits/internal/storage"
	"velocity-limits/pkg/util"
)

// GetTransactionsFromInputFile reads the input file and creates a slice of Transaction struct.
func GetTransactionsFromInputFile(config *config.Configuration, filePath string) ([]models.Transaction, error) {
	transactions := []models.Transaction{}
	inputFile, err := util.OpenFile(config, filePath)
	if err != nil {
		return nil, err
	}
	// defers input file closing so that it can be read.
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		byteValue := []byte(scanner.Text())
		var transaction models.Transaction
		// tries to unmarshal input transactions json into Transaction struct.
		if err = json.Unmarshal(byteValue, &transaction); err != nil {
			return nil, err
		}

		// append current transaction to transactions slice.
		transactions = append(transactions, transaction)
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}
	return transactions, err
}

// LoadFunds reads transactions, loads it into storage and creates a slice of Response struct.
func LoadFunds(config *config.Configuration, transactions []models.Transaction, storage *storage.Storage) []models.Response {
	responses := []models.Response{}
	for _, transaction := range transactions {
		response := ValidateAndProcessTransaction(&transaction, config, storage)
		if response != nil {
			responses = append(responses, *response)
		}
	}
	return responses
}

// Validates transaction for duplication and send it for further processing.
// Also stores processing result into responses slice.
func ValidateAndProcessTransaction(transaction *models.Transaction, config *config.Configuration, storage *storage.Storage) *models.Response {
	// Checks if load ID is repeated for the same customer ID.
	if storage.IsDuplicateTransaction(transaction.ID, transaction.CustomerID) {
		log.Printf("Ignoring a duplicate transaction: %+v\n", transaction)
		return nil
	}

	// If valid, add it to the storage and send it for processing.
	storage.AddTransaction(transaction.ID, transaction.CustomerID)
	accepted := ProcessTransaction(transaction, storage, config)
	response := models.NewResponse(transaction.ID, transaction.CustomerID, accepted)

	return response
}

// ProcessTransaction function verifies if customer account is created in the storage.
// If not it will create a new account with default velocity limits.
// If already created, then tries to reset limits based on transaction time.
// At last, it tries to load the funds from given transaction.
func ProcessTransaction(transaction *models.Transaction, storage *storage.Storage, config *config.Configuration) bool {
	// Get the account from storage
	account := storage.GetAccount(transaction.CustomerID)

	if account == nil {
		account = models.NewCustomerAccount(transaction.CustomerID)
		account.DailyLimit = models.NewDailyLimit(transaction.Time, config.MaxLoadLimitPerDay, config.MaxLoadPerDay)
		account.WeeklyLimit = models.NewWeeklyLimit(transaction.Time, config.MaxLoadLimitPerWeek)
		storage.AddAccount(account)
	} else {
		account.ResetLimits(transaction.Time, config.MaxLoadLimitPerDay, config.MaxLoadPerDay, config.MaxLoadLimitPerWeek)
	}

	return account.LoadFunds(transaction)
}

// WriteResponsesToOutputFile writes responses to the output.txt file.
// Returns any error or nil in case succeed.
func WriteResponsesToOutputFile(config *config.Configuration, responses []models.Response, filePath string) error {
	outputFile := util.CreateFile(config, filePath)
	defer outputFile.Close()
	writer := bufio.NewWriter(outputFile)

	for _, response := range responses {
		byteValue, err := json.Marshal(response)
		if err != nil {
			return err
		}

		if _, err = writer.WriteString(string(byteValue) + "\n"); err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}
