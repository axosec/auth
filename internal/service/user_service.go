package service

import (
	"context"
	"errors"

	"github.com/axosec/auth/internal/data/db"
	"github.com/axosec/auth/internal/dto"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserService struct {
	q *db.Queries
}

func NewUserService(q *db.Queries) *UserService {
	return &UserService{
		q: q,
	}
}

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (dto.User, error) {
	user, err := s.q.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.User{}, ErrUserNotFound
		}

		return dto.User{}, err
	}

	return dto.User{
		ID:        user.ID,
		Email:     user.Email,
		EmailHash: user.EmailHash,
		Username:  user.Username,

		Salt:         user.Salt,
		AuthVerifier: user.AuthVerifier,

		IdentityPublicKey:       user.IdentityPublicKey,
		EncIdentityPrivateKey:   user.EncIdentityPrivateKey,
		IdentityPrivateKeyNonce: user.IdentityPrivateKeyNonce,

		VaultPublicKey:       user.VaultPublicKey,
		EncVaultPrivateKey:   user.EncVaultPrivateKey,
		VaultPrivateKeyNonce: user.VaultPrivateKeyNonce,
	}, nil
}

func (s *UserService) LookupUser(ctx context.Context, emailHash []byte) (dto.UserLookup, error) {
	user, err := s.q.LookupUser(ctx, emailHash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.UserLookup{}, ErrUserNotFound
		}
		return dto.UserLookup{}, err
	}

	return dto.UserLookup{
		ID:                user.ID,
		Username:          user.Username,
		IdentityPublicKey: user.IdentityPublicKey,
		VaultPublicKey:    user.VaultPublicKey,
	}, nil
}

func (s *UserService) LookupUsers(ctx context.Context, ids []uuid.UUID) ([]dto.UserLookup, error) {
	users, err := s.q.LookupUsers(ctx, ids)
	if err != nil {
		return nil, err
	}

	res := make([]dto.UserLookup, len(users))
	for i, u := range users {
		res[i] = dto.UserLookup{
			ID:                u.ID,
			Username:          u.Username,
			IdentityPublicKey: u.IdentityPublicKey,
			VaultPublicKey:    u.VaultPublicKey,
		}
	}

	return res, nil
}
