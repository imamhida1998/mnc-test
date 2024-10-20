package request

type Payment struct {
	Amount  int    `json:"amount"`
	Remarks string `json:"remarks"`
}
