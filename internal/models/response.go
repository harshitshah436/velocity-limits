package models

type Response struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	Accepted   bool   `json:"accepted"`
}

func NewResponse(id string, custID string, accepted bool) *Response {
	return &Response{
		ID:         id,
		CustomerID: custID,
		Accepted:   accepted,
	}
}
