package dto

type TransactionRequest struct {
	AccountId       int     `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
}

func (r TransactionRequest) ValidateAmount() {

}
