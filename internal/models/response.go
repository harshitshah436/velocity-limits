// Package models represents model structs and its functions.
package models

// Response struct stores load ID, customer ID and accepted flag.
// Accepted flag represents transaction was load successfully or failed.
type Response struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}

// Returns a new Response struct.
func NewResponse(id string, custID string, accepted bool) *Response {
	return &Response{
		ID:         id,
		CustomerID: custID,
		Accepted:   accepted,
	}
}
