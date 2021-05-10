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
func GetTransactionsFromInputFile(config *config.Configuration) ([]models.Transaction, error) {
	transactions := []models.Transaction{}
	inputFile, err := util.OpenFile(config)
	if err != nil {
		return nil, err
	}
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		byteValue := []byte(scanner.Text())
		var transaction models.Transaction
		if err = json.Unmarshal(byteValue, &transaction); err != nil {
			return nil, err
		}

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

func ValidateAndProcessTransaction(transaction *models.Transaction, config *config.Configuration, storage *storage.Storage) *models.Response {
	// Checks if load ID is repeated for the same customer ID
	if storage.IsDuplicateTransaction(transaction.ID, transaction.CustomerID) {
		log.Printf("Ignoring a duplicate transaction: %+v\n", transaction)
		return nil
	}

	storage.AddTransaction(transaction.ID, transaction.CustomerID)
	accepted := ProcessTransaction(transaction, storage, config)
	response := models.NewResponse(transaction.ID, transaction.CustomerID, accepted)

	return response
}

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
func WriteResponsesToOutputFile(config *config.Configuration, responses []models.Response) error {
	outputFile := util.CreateFile(config)
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
