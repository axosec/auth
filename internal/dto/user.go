package dto

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `json:"id" binding:"required,uuid4"`
	Email     string    `json:"email" binding:"required,email,min=6,max=250"`
	EmailHash []byte    `json:"email_hash" binding:"required,min=32"`
	Username  string    `json:"username" binding:"required,min=6,max=250"`

	Salt         []byte `json:"salt" binding:"required,min=32"`
	AuthVerifier []byte `json:"auth_verifier" binding:"required,min=32"`

	IdentityPublicKey       []byte `json:"identity_public_key" binding:"required,min=32"`
	EncIdentityPrivateKey   []byte `json:"enc_identity_private_key" binding:"required"`
	IdentityPrivateKeyNonce []byte `json:"identity_private_key_nonce" binding:"required,min=24"`

	VaultPublicKey       []byte `json:"vault_public_key" binding:"required,min=32"`
	EncVaultPrivateKey   []byte `json:"enc_vault_private_key" binding:"required"`
	VaultPrivateKeyNonce []byte `json:"vault_private_key_nonce" binding:"required,min=24"`
}
