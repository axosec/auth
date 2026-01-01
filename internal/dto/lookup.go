package dto

import "github.com/google/uuid"

type UserLookup struct {
	ID                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
	IdentityPublicKey []byte    `json:"identity_public_key"`
	VaultPublicKey    []byte    `json:"vault_public_key"`
}

type LookupUserRequest struct {
	EmailHash []byte `json:"email_hash" binding:"required"`
}

type LookupUsersRequest struct {
	IDs []uuid.UUID `json:"ids" binding:"required"`
}
