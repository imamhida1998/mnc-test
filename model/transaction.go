package model

import "time"

type Transaction struct {
	TransactionId   string    `json:"transaction_id"`
	PaymentId       string    `json:"payment_id,omitempty"`
	TopUpID         string    `json:"top_up_id,omitempty"`
	TransferId      string    `json:"transfer_id,omitempty"`
	Status          string    `json:"status"`
	UserID          string    `json:"user_id"`
	TransactionType string    `json:"transaction_type"`
	Amount          int       `json:"amount"`
	Remarks         string    `json:"remarks,omitempty"`
	BalanceBefore   int       `json:"balance_before"`
	BalanceAfter    int       `json:"balance_after"`
	CreatedAt       time.Time `json:"created_date"`
}
