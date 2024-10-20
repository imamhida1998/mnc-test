package response

type Registration struct {
	UserID      string `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	CreatedDate string `json:"created_date"`
}
