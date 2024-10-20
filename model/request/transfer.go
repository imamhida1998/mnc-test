package request

type Transfer struct {
	TargetUser string `json:"target_user"`
	Amount     int    `json:"amount"`
	Remarks    string `json:"remarks"`
}
