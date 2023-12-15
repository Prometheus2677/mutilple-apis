package dto

type AddBalanceRequest struct {
	Amount int64 `json:"amount" form:"amount" binding:"required,numeric,min=10000"`
}
