package dto

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email,min=5,max=250"`
	Username string `json:"username" binding:"required,min=6,max=250"`

	Salt          []byte `json:"salt" binding:"required,min=32"`
	AuthVerifier  []byte `json:"auth_verifier" binding:"required,min=32"`
	PublicKey     []byte `json:"public_key" binding:"required,min=32"`
	EncPrivateKey []byte `json:"enc_private_key" binding:"required"`
}
