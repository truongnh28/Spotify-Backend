package dto

type LoginRequest struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

type LoginResponseData struct {
	Account     *Account `json:"account,omitempty"`
	AccessToken string   `json:"access_token,omitempty"`
}

type Account struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Code     string `json:"code,omitempty"`
	Role     string `json:"role,omitempty"`
	Status   string `json:"status,omitempty"`
}
