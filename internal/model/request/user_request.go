package request

// RegisterRequest represents the user registration request
// @Description User registration request
type RegisterRequest struct {
	Username string `json:"username" example:"john_doe" binding:"required"`
	Password string `json:"password" example:"secure123" binding:"required"`
	FullName string `json:"full_name" example:"John Doe" binding:"required"`
	Email    string `json:"email" example:"john@example.com" binding:"required,email"`
	Role     string `json:"role" example:"employee" binding:"required,oneof=admin employee"`
}

// LoginRequest represents the login credentials
// @Description Login request
type LoginRequest struct {
	Username string `json:"username" example:"john_doe" binding:"required"`
	Password string `json:"password" example:"secure123" binding:"required"`
}

// UpdateUserRequest represents the user update request
// @Description User update request
type UpdateUserRequest struct {
	Username string `json:"username" example:"john_doe" binding:"required"`
	FullName string `json:"full_name" example:"John Doe Updated" binding:"required"`
	Email    string `json:"email" example:"john.updated@example.com" binding:"required,email"`
	Role     string `json:"role" example:"employee" binding:"required,oneof=admin employee"`
}
