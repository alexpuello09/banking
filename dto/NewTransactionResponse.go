package dto

type TransactionResponse struct {
	TransactionId string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}
