package service

import (
	"testing"
	"velocity-limits/config"
	"velocity-limits/internal/models"
	"velocity-limits/internal/service"
	"velocity-limits/internal/storage"

	"github.com/stretchr/testify/assert"
)

var configVar = config.LoadConfig("../../../config/")
var storageVar = storage.NewStorage()
var transactionsVar []models.Transaction
var responsesVar []models.Response

func TestGetTransactionsFromInputFile(t *testing.T) {
	t.Run("should read the file and get transactions", func(t *testing.T) {
		transactions, err := service.GetTransactionsFromInputFile(&configVar, "../../../")
		assert.NoError(t, err)
		assert.NotZero(t, transactions)
		transactionsVar = transactions
	})
}

func TestValidateAndProcessTransaction(t *testing.T) {
	t.Run("should validate transactions and process them to create responses", func(t *testing.T) {
		responses := service.LoadFunds(&configVar, transactionsVar, storageVar)
		expectedType := []models.Response{}
		assert.NotZero(t, responses)
		assert.IsType(t, expectedType, responses)
		responsesVar = responses
	})
}

func TestWriteResponsesToOutputFile(t *testing.T) {
	t.Run("should write responses to the output file", func(t *testing.T) {
		err := service.WriteResponsesToOutputFile(&configVar, responsesVar, "../../../")
		assert.NoError(t, err)
	})
}
