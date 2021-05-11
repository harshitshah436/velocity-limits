package models

import (
	"testing"
	"time"

	"velocity-limits/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestGetParsedAmount(t *testing.T) {
	t.Run("parses input load amount correcly", func(t *testing.T) {
		parsedTime, _ := time.Parse(time.RFC3339, "2011-01-11T06:08:12Z")
		transaction := &models.Transaction{
			ID:         "123",
			CustomerID: "1234",
			Amount:     "$11000",
			Time:       parsedTime,
		}
		result := transaction.GetParsedAmount()
		expected := float64(11000)

		assert.IsType(t, expected, result)
		assert.Equal(t, expected, result)
	})
}
