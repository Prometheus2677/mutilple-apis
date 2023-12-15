package dto

// Define form validation when user hit /register API
type RegisterRequest struct {
	Username string `json:"username" form:"username" binding:"required,lowercase"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
