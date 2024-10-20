package model

import "time"

type Transfer struct {
	TransferId string    `json:"transfer_id"`
	UserId     string    `json:"user_id"`
	Amount     int       `json:"amount"`
	Remarks    string    `json:"remarks"`
	CreatedAt  time.Time `json:"created_at"`
}
