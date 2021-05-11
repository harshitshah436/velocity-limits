package models

import (
	"testing"
	"time"

	"velocity-limits/internal/models"
	"velocity-limits/pkg/util"

	"github.com/stretchr/testify/assert"
)

func TestNewCustomerAccount(t *testing.T) {
	t.Run("returns a new customer account", func(t *testing.T) {
		expectedCustomerAccount := &models.CustomerAccount{
			CustomerID: "1234",
		}
		result := models.NewCustomerAccount("1234")
		assert.Equal(t, expectedCustomerAccount, result)
	})
}

func TestNewDailyLimit(t *testing.T) {
	t.Run("returns a new velocity limits per day", func(t *testing.T) {
		expectedDailyLimit := &models.DailyLimit{
			Date:         util.GetBeginningOfTheDay(time.Now()),
			MaxLoadLimit: 5000,
			MaxLoad:      3,
		}
		result := models.NewDailyLimit(time.Now(), 5000, 3)
		assert.Equal(t, expectedDailyLimit, result)
	})
}

func TestNewWeeklyLimit(t *testing.T) {
	t.Run("returns a new velocity limits per week", func(t *testing.T) {
		expectedWeeklyLimit := &models.WeeklyLimit{
			Date:         util.GetBeginningOfTheWeek(time.Now()),
			MaxLoadLimit: 20000,
		}
		result := models.NewWeeklyLimit(time.Now(), 20000)
		assert.Equal(t, expectedWeeklyLimit, result)
	})
}

func TestValidateDailyLimit(t *testing.T) {
	t.Run("returns true when loading below max load limit per day", func(t *testing.T) {
		dailyLimit := models.NewDailyLimit(time.Now(), 5000, 3)
		valid := dailyLimit.Validate(3000)
		assert.True(t, valid)
	})
	t.Run("returns true when loading exactly same max load limit per day", func(t *testing.T) {
		dailyLimit := models.NewDailyLimit(time.Now(), 5000, 3)
		valid := dailyLimit.Validate(5000)
		assert.True(t, valid)
	})
	t.Run("returns false when loading more than allowed max load limit per day", func(t *testing.T) {
		dailyLimit := models.NewDailyLimit(time.Now(), 5000, 3)
		valid := dailyLimit.Validate(5001)
		assert.False(t, valid)
	})
}

func TestValidateWeeklyLimit(t *testing.T) {
	t.Run("returns true when loading below max load limit per week", func(t *testing.T) {
		WeeklyLimit := models.NewWeeklyLimit(time.Now(), 20000)
		valid := WeeklyLimit.Validate(2000)
		assert.True(t, valid)
	})
	t.Run("returns true when loading exactly same max load limit per week", func(t *testing.T) {
		WeeklyLimit := models.NewWeeklyLimit(time.Now(), 20000)
		valid := WeeklyLimit.Validate(20000)
		assert.True(t, valid)
	})
	t.Run("returns false when loading more than allowed max load limit per week", func(t *testing.T) {
		WeeklyLimit := models.NewWeeklyLimit(time.Now(), 20000)
		valid := WeeklyLimit.Validate(21000)
		assert.False(t, valid)
	})
}

func TestDailyUpdateLimits(t *testing.T) {
	t.Run("updates allocated daily limit and reflecting amount will be reduced", func(t *testing.T) {
		dailyLimit := models.NewDailyLimit(time.Now(), 5000, 3)
		dailyLimit.UpdateLimits(2000)
		assert.Equal(t, float64(3000), dailyLimit.MaxLoadLimit)
	})
}

func TestWeeklyUpdateLimits(t *testing.T) {
	t.Run("updates allocated weekly limit and reflecting amount will be reduced", func(t *testing.T) {
		weeklyLimit := models.NewWeeklyLimit(time.Now(), 20000)
		weeklyLimit.UpdateLimits(15000)
		assert.Equal(t, float64(5000), weeklyLimit.MaxLoadLimit)
	})
}

func TestResetLimits(t *testing.T) {
	t.Run("should not reset velocity limits if load time is within daily/weekly limits", func(t *testing.T) {
		customerAccount := models.NewCustomerAccount("1234")
		now := time.Now()
		customerAccount.DailyLimit = models.NewDailyLimit(now, 5000, 3)
		customerAccount.WeeklyLimit = models.NewWeeklyLimit(now, 20000)

		customerAccount.ResetLimits(time.Now(), 1000, 5, 10000)

		assert.Equal(t, float64(5000), customerAccount.DailyLimit.MaxLoadLimit)
		assert.Equal(t, 3, customerAccount.DailyLimit.MaxLoad)
		assert.Equal(t, float64(20000), customerAccount.WeeklyLimit.MaxLoadLimit)
	})

	t.Run("should reset velocity limits if load time is before current day/week", func(t *testing.T) {
		customerAccount := models.NewCustomerAccount("1234")
		previousMonth := time.Now().AddDate(0, -1, 0)
		customerAccount.DailyLimit = models.NewDailyLimit(previousMonth, 5000, 3)
		customerAccount.WeeklyLimit = models.NewWeeklyLimit(previousMonth, 20000)

		now := time.Now()
		customerAccount.ResetLimits(now, 1000, 5, 10000)

		assert.Equal(t, float64(1000), customerAccount.DailyLimit.MaxLoadLimit)
		assert.Equal(t, 5, customerAccount.DailyLimit.MaxLoad)
		assert.Equal(t, float64(10000), customerAccount.WeeklyLimit.MaxLoadLimit)
	})
}

func TestLoadFunds(t *testing.T) {

	t.Run("should return true when max load per day, max load per week and max load limits are not reached", func(t *testing.T) {
		customerAccount := models.NewCustomerAccount("1234")
		now := time.Now()
		customerAccount.DailyLimit = models.NewDailyLimit(now, 5000, 3)
		customerAccount.WeeklyLimit = models.NewWeeklyLimit(now, 20000)

		txn := models.Transaction{
			ID:         "123",
			CustomerID: "1234",
			Amount:     "$3000",
			Time:       now,
		}

		success := customerAccount.LoadFunds(&txn)
		assert.True(t, success)
		assert.Equal(t, float64(3000), customerAccount.Balance)
		assert.Equal(t, float64(2000), customerAccount.DailyLimit.MaxLoadLimit)
		assert.Equal(t, float64(17000), customerAccount.WeeklyLimit.MaxLoadLimit)
		assert.Equal(t, 2, customerAccount.DailyLimit.MaxLoad)
	})

	t.Run("should return false when max load per day, max load per week and max load limits are reached", func(t *testing.T) {
		customerAccount := models.NewCustomerAccount("1234")
		now := time.Now()
		customerAccount.DailyLimit = models.NewDailyLimit(now, 5000, 3)
		customerAccount.WeeklyLimit = models.NewWeeklyLimit(now, 20000)

		txn := models.Transaction{
			ID:         "123",
			CustomerID: "1234",
			Amount:     "$8000",
			Time:       now,
		}

		result := customerAccount.LoadFunds(&txn)
		assert.False(t, result)
		assert.Equal(t, float64(0), customerAccount.Balance)
		assert.Equal(t, float64(5000), customerAccount.DailyLimit.MaxLoadLimit)
		assert.Equal(t, float64(20000), customerAccount.WeeklyLimit.MaxLoadLimit)
		assert.Equal(t, 3, customerAccount.DailyLimit.MaxLoad)
	})
}
