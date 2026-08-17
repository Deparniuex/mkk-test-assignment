package auth

import "errors"

var (
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token is expired")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
