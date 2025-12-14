package dto

type InitLoginRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type InitLoginResponse struct {
	Salt []byte `json:"salt"`
}

type LoginRequest struct {
	Email        string `json:"email" binding:"required,email"`
	AuthVerifier []byte `json:"auth_verifier" binding:"required,min=32"`
}
