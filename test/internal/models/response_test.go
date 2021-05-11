package models

import (
	"testing"

	"velocity-limits/internal/models"

	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {
	t.Run("returns expected response", func(t *testing.T) {
		expectedResponse := &models.Response{
			ID:         "123",
			CustomerID: "1234",
			Accepted:   false,
		}
		actualResponse := models.NewResponse("123", "1234", false)
		assert.Equal(t, expectedResponse, actualResponse)
	})
}
