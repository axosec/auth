package service

import (
	"context"
	"crypto/sha512"
	"errors"

	"github.com/axosec/auth/internal/data/db"
	"github.com/axosec/auth/internal/dto"
	"github.com/axosec/core/crypto/token"
	"github.com/jackc/pgx/v5/pgconn"
)

type AuthService struct {
	q   *db.Queries
	jwt *token.JWTManager
}

func NewAuthService(q *db.Queries, m *token.JWTManager) *AuthService {
	return &AuthService{
		q:   q,
		jwt: m,
	}
}

var (
	ErrUserAlreadyExists = errors.New("email already currently in use")
	ErrInvalidKeyLength  = errors.New("cryptographic keys must be correct length")
)

func (s *AuthService) RegisterUser(ctx context.Context, req dto.RegisterRequest) error {
	if len(req.PublicKey) != 32 {
		return ErrInvalidKeyLength
	}

	if len(req.Salt) != 32 {
		return ErrInvalidKeyLength
	}
	verifierHash := sha512.Sum512(req.AuthVerifier)

	verifierSlice := verifierHash[:]

	args := db.CreateUserParams{
		Email:         req.Email,
		Salt:          req.Salt,
		AuthVerifier:  verifierSlice,
		PublicKey:     req.PublicKey,
		EncPrivateKey: req.EncPrivateKey,
	}

	_, err := s.q.CreateUser(ctx, args)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return ErrUserAlreadyExists
			}
		}
		return err
	}

	return nil
}
