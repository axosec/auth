package service

import (
	"bytes"
	"context"
	"crypto/sha512"
	"errors"
	"time"

	"github.com/axosec/auth/internal/data/db"
	"github.com/axosec/auth/internal/dto"
	"github.com/axosec/core/crypto/token"
	"github.com/axosec/core/utils"
	"github.com/jackc/pgx/v5"
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
	ErrUserAlreadyExists  = errors.New("email already currently in use")
	ErrInvalidKeyLength   = errors.New("cryptographic keys must be correct length")
	ErrUserNotFound       = errors.New("user not found")
	ErrInvalidCredentials = errors.New("invalid user credentials")
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
		Username:      req.Username,
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

func (s *AuthService) InitLogin(ctx context.Context, req dto.InitLoginRequest) (dto.InitLoginResponse, error) {
	salt, err := s.q.GetSaltByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			fakeSalt, saltErr := utils.GenerateSalt(32)
			if saltErr != nil {
				return dto.InitLoginResponse{}, saltErr
			}
			return dto.InitLoginResponse{Salt: fakeSalt}, nil
		}

		return dto.InitLoginResponse{}, err
	}

	return dto.InitLoginResponse{Salt: salt}, nil
}

func (s *AuthService) Login(ctx context.Context, req dto.LoginRequest) (dto.User, string, error) {
	user, err := s.q.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return dto.User{}, "", ErrUserNotFound
		}

		return dto.User{}, "", err
	}

	userHash := sha512.Sum512(req.AuthVerifier)

	if !bytes.Equal(userHash[:], user.AuthVerifier) {
		return dto.User{}, "", ErrInvalidCredentials
	}

	token, err := s.jwt.Issue(user.ID.String(), time.Hour*24)
	if err != nil {
		return dto.User{}, "", err
	}

	return dto.User{
		ID: user.ID,
		Email: user.Email,
		Username: user.Username,
		Salt: user.Salt,
		AuthVerifier: user.AuthVerifier,
		PublicKey: user.PublicKey,
		EncPrivateKey: user.EncPrivateKey,
	}, token, err
}
