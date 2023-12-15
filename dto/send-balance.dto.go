package dto

type SendBalanceRequest struct {
	Amount    int64  `json:"amount" form:"amount" binding:"required,numeric,min=10000"`
	Recipient string `json:"recipient" form:"recipient" binding:"required"`
}
