package auth

import (
	"context"
	"errors"
	"portal_back/pkg/app/token"
)

var ErrUserNotFound = errors.New("user not found")

type LoginData struct {
	Login    string
	Password string
}

type Service interface {
	GetSaltByLogin(ctx context.Context, login string) (string, error)
	Login(ctx context.Context, logData LoginData) (token.Tokens, error)
}

func NewService(repository Repository, tokenService token.Service) Service {
	return &service{repository: repository, tokenService: tokenService}
}

type service struct {
	repository   Repository
	tokenService token.Service
}

func (s *service) GetSaltByLogin(ctx context.Context, login string) (string, error) {
	return s.repository.GetSaltByLogin(ctx, login)
}

func (s *service) Login(ctx context.Context, logData LoginData) (token.Tokens, error) {
	password, err := s.repository.GetPasswordByLogin(ctx, logData.Login)
	if err != nil {
		return token.Tokens{}, err
	}
	if password != logData.Password {
		return token.Tokens{}, ErrUserNotFound
	}

	userID, err := s.repository.GetUserIDByLogin(ctx, logData.Login)
	if err != nil {
		return token.Tokens{}, err
	}

	return s.tokenService.GenerateTokensForUser(ctx, userID)
}
