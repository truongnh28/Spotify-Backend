package dto

type CreateAccountRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role" binding:"required,oneof=SuperAdmin Admin User"`
	Status   string `json:"status" binding:"required,oneof=Active Blocked"`
}

type UpdateAccountRequest struct {
	Role   string `json:"role" binding:"oneof=SuperAdmin Admin User"`
	Status string `json:"status" binding:"oneof=Active Blocked"`
}

type AccountsResponse struct {
	Accounts []Account `json:"accounts"`
	StatusError
}

type AccountResponse struct {
	Account Account `json:"account"`
	StatusError
}
