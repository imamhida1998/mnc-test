package response

type TopUp struct {
	TopUpID       string `json:"top_up_id"`
	AmountTopUp   int    `json:"amount_top_up"`
	BalanceBefore int    `json:"balance_before"`
	BalanceAfter  int    `json:"balance_after"`
	CreatedDate   string `json:"created_date"`
}
