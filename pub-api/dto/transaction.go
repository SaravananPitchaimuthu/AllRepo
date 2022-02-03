package dto

import "github.com/SaravananPitchaimuthu/pub-api/errs"

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionType string  `json:"transaction_type"`
	TransactionDate string  `json:"transaction_date"`
	CustomerId      string  `json:"customer_id"`
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return r.TransactionType == "withdrawal"
}

func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	return r.TransactionType == "deposit"
}

func (r TransactionRequest) Validate() *errs.AppError {
	if !r.IsTransactionTypeDeposit() && r.IsTransactionTypeWithdrawal() {
		return errs.NewValidationError("Transaction type either Deposit or Withdrawal")
	}

	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	TransactionDate string  `json:"transaction_date"`
	TransactionType string  `json:"transaction_type"`
}
