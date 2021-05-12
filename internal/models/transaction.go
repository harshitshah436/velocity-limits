// Package models represents model structs and its functions.
package models

import (
	"log"
	"strconv"
	"strings"
	"time"
)

// Transaction struct stores load ID, customer ID, load amount and transaction time.
// It represents the transaction payload from the input file.
type Transaction struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	Amount     string    `json:"load_amount"`
	Time       time.Time `json:"time"`
}

// GetParsedAmount function parses amount from the transaction struct
// into float64 type and removes $ sign.
func (txn *Transaction) GetParsedAmount() float64 {
	parsedAmount, err := strconv.ParseFloat(strings.Trim(txn.Amount, "$"), 64)
	if err != nil {
		log.Fatal("Error - GetParsedAmount function error: ", err)
	}
	return parsedAmount
}
