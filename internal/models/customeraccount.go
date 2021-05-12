// Package models represents model structs and its functions.
package models

import (
	"time"
	"velocity-limits/pkg/util"
)

// CustomerAccount struct stores customer balance and current velocity limits.
type CustomerAccount struct {
	CustomerID  string
	Balance     float64
	DailyLimit  *DailyLimit
	WeeklyLimit *WeeklyLimit
}

// DailyLimit struct represents daily load limits and max loads.
type DailyLimit struct {
	Date         time.Time
	MaxLoadLimit float64
	MaxLoad      int
}

// WeeklyLimit struct represents weekly load limits.
type WeeklyLimit struct {
	Date         time.Time
	MaxLoadLimit float64
}

// Returns a new customer account struct.
func NewCustomerAccount(customerID string) *CustomerAccount {
	return &CustomerAccount{
		CustomerID: customerID,
	}
}

// Returns a new daily limit struct.
func NewDailyLimit(d time.Time, maxLoadLimit float64, maxLoad int) *DailyLimit {
	return &DailyLimit{
		Date:         util.GetBeginningOfTheDay(d),
		MaxLoadLimit: maxLoadLimit,
		MaxLoad:      maxLoad,
	}
}

// Returns a new weekly limit struct.
func NewWeeklyLimit(d time.Time, maxLoadLimit float64) *WeeklyLimit {
	return &WeeklyLimit{
		Date:         util.GetBeginningOfTheWeek(d),
		MaxLoadLimit: maxLoadLimit,
	}
}

// Validates daily velocity limits are not reached
func (dl *DailyLimit) Validate(amount float64) bool {
	return dl.MaxLoadLimit-amount >= 0 && dl.MaxLoad-1 >= 0
}

// Updates daily limit struct.
func (dl *DailyLimit) UpdateLimits(amount float64) {
	dl.MaxLoadLimit -= amount
	dl.MaxLoad--
}

// Validates daily velocity limits are not reached
func (wl *WeeklyLimit) Validate(amount float64) bool {
	return wl.MaxLoadLimit-amount >= 0
}

// Updates weekly limit struct.
func (wl *WeeklyLimit) UpdateLimits(amount float64) {
	wl.MaxLoadLimit -= amount
}

// Reset daily limits and/or weekly limits depending on a transaction time.
func (c *CustomerAccount) ResetLimits(transactionTime time.Time, maxLoadLimitPerDay float64, maxLoad int, maxLoadLimitPerWeek float64) {
	transactionDay := util.GetBeginningOfTheDay(transactionTime)
	if transactionDay.After(c.DailyLimit.Date) {
		c.DailyLimit.Date = transactionDay
		c.DailyLimit.MaxLoadLimit = maxLoadLimitPerDay
		c.DailyLimit.MaxLoad = maxLoad
	}
	transactionWeek := util.GetBeginningOfTheWeek(transactionTime)
	if transactionWeek.After(c.WeeklyLimit.Date) {
		c.WeeklyLimit.Date = transactionWeek
		c.WeeklyLimit.MaxLoadLimit = maxLoadLimitPerWeek
	}
}

// Tries to load fund if it's within daily and weekly velocity limits.
func (c *CustomerAccount) LoadFunds(txn *Transaction) bool {
	if !c.DailyLimit.Validate(txn.GetParsedAmount()) {
		return false
	}

	if !c.WeeklyLimit.Validate(txn.GetParsedAmount()) {
		return false
	}

	c.Balance += txn.GetParsedAmount()
	c.DailyLimit.UpdateLimits(txn.GetParsedAmount())
	c.WeeklyLimit.UpdateLimits(txn.GetParsedAmount())
	return true
}
