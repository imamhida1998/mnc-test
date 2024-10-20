package model

import "time"

type User struct {
	UserId      string    `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	Pin         string    `json:"pin,omitempty"`
	Balance     int       `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
