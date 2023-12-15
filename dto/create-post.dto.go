package dto

type CreateRequest struct {
	Title string `json:"title" form:"title" binding:"required"`
	Body  string `json:"body" form:"body" binding:"required"`
}
