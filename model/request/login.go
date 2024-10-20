package request

type Login struct {
	PhoneNumber string `json:"phone_number"`
	Pin         string `json:"pin"`
}
