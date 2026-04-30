package auth

type LoginRequest struct {
	Username string `json:"username"`
	Code     string `json:"code"`
}
