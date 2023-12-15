package dto

// Define form validation when user hit /login API
type LoginCredential struct {
	Account  string `json:"account" form:"account" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
