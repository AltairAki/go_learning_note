package optimization

// Request is the input of the filter
type Request struct {
	TransactionID string `json:"transaction_id"`
	PayLoad       []int  `json:"payload"`
}

// Response is the output of the filter
type Response struct {
	TransactionID string `json:"transaction_id"`
	Expression    string `json:"exp"`
}
