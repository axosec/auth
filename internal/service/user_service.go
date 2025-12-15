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
		ID:            user.ID,
		Email:         user.Email,
		Username:      user.Username,
		Salt:          user.Salt,
		AuthVerifier:  user.AuthVerifier,
		PublicKey:     user.PublicKey,
		EncPrivateKey: user.EncPrivateKey,
	}, nil
}
