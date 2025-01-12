package services

import (
	"errors"
)

type AuthService struct{}

func NewAuthService() *AuthService {
	return new(AuthService)
}

func (s *AuthService) Login(username, password string) (string, error) {
	if username == "admin" && password == "123" {
		return "fake-jwt-token-12345", nil
	}
	return "", errors.New("username or password error")
}
