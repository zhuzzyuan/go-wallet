package models

type DepositRequest struct {
	Email string `json:"email" binding:"required"`
	Chain string `json:"chain" binding:"required"`
}

type WithdrawRequest struct {
	Email       string `json:"email" binding:"required"`
	CoinType    string `json:"coin_type" binding:"required"`
	Chain       string `json:"chain" binding:"required"`
	Value       string `json:"value" binding:"required"`
	Destination string `json:"destination" binding:"required"`
}

type TransferRequest struct {
	Email            string `json:"email" binding:"required"`
	CoinType         string `json:"coin_type" binding:"required"`
	Chain            string `json:"chain" binding:"required"`
	Value            string `json:"value" binding:"required"`
	DestinationEmail string `json:"destination_email" binding:"required"`
}

type BalanceRequest struct {
	Email string `form:"email" binding:"required"`
}

type BalanceResponse struct {
	CoinType string `json:"coin_type"`
	Chain    string `json:"chain"`
	Value    string `json:"value"`
}

type TxHistoryResponse struct {
	From      string `json:"from"`
	To        string `json:"to"`
	Value     string `json:"value"`
	Chain     string `json:"chain"`
	CoinType  string `json:"coin_type"`
	Timestamp uint64 `json:"timestamp"`
}

type GetTokenRequest struct {
	Email string `json:"email" binding:"required"`
}
