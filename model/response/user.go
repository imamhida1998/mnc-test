package response

import "time"

type User struct {
	UserId      string    `json:"user_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Address     string    `json:"address"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
