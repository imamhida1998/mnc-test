package response

type Transfer struct {
	TransferId    string `json:"transfer_id"`
	Amount        int    `json:"amount"`
	Remarks       string `json:"remarks"`
	BalanceBefore int    `json:"balance_before"`
	BalanceAfter  int    `json:"balance_after"`
	CreatedDate   string `json:"created_date"`
}
