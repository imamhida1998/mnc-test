package response

type Payment struct {
	TopUpID       string `json:"top_up_id"`
	Amount        int    `json:"amount"`
	Remarks       string `json:"remarks"`
	BalanceBefore int    `json:"balance_before"`
	BalanceAfter  int    `json:"balance_after"`
	CreatedDate   string `json:"created_date"`
}
