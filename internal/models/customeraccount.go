package models

import (
	"time"
	"velocity-limits/pkg/util"
)

type CustomerAccount struct {
	CustomerID  string
	Balance     float64
	DailyLimit  *DailyLimit
	WeeklyLimit *WeeklyLimit
}

type DailyLimit struct {
	Date         time.Time
	MaxLoadLimit float64
	MaxLoad      int
}

type WeeklyLimit struct {
	Date         time.Time
	MaxLoadLimit float64
}

func NewCustomerAccount(customerID string) *CustomerAccount {
	return &CustomerAccount{
		CustomerID: customerID,
	}
}

func NewDailyLimit(d time.Time, maxLoadLimit float64, maxLoad int) *DailyLimit {
	return &DailyLimit{
		Date:         util.GetBeginningOfTheDay(d),
		MaxLoadLimit: maxLoadLimit,
		MaxLoad:      maxLoad,
	}
}

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

func (dl *DailyLimit) UpdateLimits(amount float64) {
	dl.MaxLoadLimit -= amount
	dl.MaxLoad--
}

// Validates daily velocity limits are not reached
func (wl *WeeklyLimit) Validate(amount float64) bool {
	return wl.MaxLoadLimit-amount >= 0
}

func (wl *WeeklyLimit) UpdateLimits(amount float64) {
	wl.MaxLoadLimit -= amount
}

func (c *CustomerAccount) ResetLimits(t time.Time, maxLoadLimitPerDay float64, maxLoad int, maxLoadLimitPerWeek float64) {
	transactionDay := util.GetBeginningOfTheDay(t)
	if transactionDay.After(c.DailyLimit.Date) {
		c.DailyLimit.Date = transactionDay
		c.DailyLimit.MaxLoadLimit = maxLoadLimitPerDay
		c.DailyLimit.MaxLoad = maxLoad
	}
	transactionWeek := util.GetBeginningOfTheWeek(t)
	if transactionWeek.After(c.WeeklyLimit.Date) {
		c.WeeklyLimit.Date = transactionWeek
		c.WeeklyLimit.MaxLoadLimit = maxLoadLimitPerWeek
	}
}

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
