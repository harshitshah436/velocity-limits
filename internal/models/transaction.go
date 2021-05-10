package models

import (
	"log"
	"strconv"
	"strings"
	"time"
)

type Transaction struct {
	ID         string    `json:"id"`
	CustomerID string    `json:"customer_id"`
	Amount     string    `json:"load_amount"`
	Time       time.Time `json:"time"`
}

func (txn *Transaction) GetParsedAmount() float64 {
	parsedAmount, err := strconv.ParseFloat(strings.Trim(txn.Amount, "$"), 64)
	if err != nil {
		log.Fatal("Error - GetParsedAmount function error: ", err)
	}
	return parsedAmount
}
