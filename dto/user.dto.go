package dto

// Define validation when update/store user data to the database
type UserDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Username string `json:"username" form:"username" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}
