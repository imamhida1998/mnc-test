package model

import "time"

type Payment struct {
	PaymentId string    `json:"paymentId"`
	UserId    string    `json:"user_id"`
	Amount    int       `json:"amount"`
	Remarks   string    `json:"remarks"`
	CreatedAt time.Time `json:"created_at"`
}
