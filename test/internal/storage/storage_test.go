package storage

import (
	"testing"

	"velocity-limits/internal/models"
	"velocity-limits/internal/storage"

	"github.com/stretchr/testify/assert"
)

func TestNewStorage(t *testing.T) {
	t.Run("should return a new storage", func(t *testing.T) {
		newStorage := storage.NewStorage()
		assert.NotNil(t, newStorage)
	})
}

func TestStorageFunctions(t *testing.T) {
	newStorage := storage.NewStorage()

	t.Run("should get a nil when no account is added to the storage", func(t *testing.T) {
		result := newStorage.GetAccount("1234")
		assert.Nil(t, result)
	})

	t.Run("should add an account to the storage", func(t *testing.T) {
		customerAccount := &models.CustomerAccount{CustomerID: "1234"}
		newStorage.AddAccount(customerAccount)
		assert.Equal(t, newStorage.GetAccount("1234"), customerAccount)
	})

	t.Run("should get an added account from the storage", func(t *testing.T) {
		expectedAccount := &models.CustomerAccount{CustomerID: "1234"}
		assert.Equal(t, newStorage.GetAccount("1234"), expectedAccount)
	})

	t.Run("should add a tranasaction to the storage struct", func(t *testing.T) {
		newStorage.AddTransaction("123", "2345")
		bool := newStorage.IsDuplicateTransaction("123", "2345")
		assert.Equal(t, bool, true)
	})

	t.Run("should return true when adding a duplicate transaction", func(t *testing.T) {
		bool := newStorage.IsDuplicateTransaction("123", "2345")
		assert.Equal(t, bool, true)
	})

	t.Run("should return false when adding a new transaction", func(t *testing.T) {
		bool := newStorage.IsDuplicateTransaction("123", "3456")
		assert.Equal(t, bool, false)
	})
}
