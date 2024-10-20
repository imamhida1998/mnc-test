package model

import "time"

type TopUp struct {
	TopUpId     string    `json:"top_up_id"`
	UserId      string    `json:"user_id"`
	AmountTopUp int       `json:"amount_top_up"`
	CreatedAt   time.Time `json:"created_at"`
}
