package impl

import (
	"context"
	"time"
	"tracker/internal/auth"
	"tracker/internal/user"
)

type authImpl struct {
	userRepository user.Repository
	jwtSecret      []byte
	tokenTTL       time.Duration
}

func NewAuthUC(user user.Repository, jwtSecret []byte, tokenTTL time.Duration) auth.AuthUC {
	return &authImpl{userRepository: user, jwtSecret: jwtSecret, tokenTTL: tokenTTL}
}

func (a *authImpl) Authenticate(ctx context.Context, email string, password string) (string, error) {
	user, err := a.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}
	err = user.ComparePassword(password)
	if err != nil {
		return "", auth.ErrInvalidCredentials
	}
	token, err := auth.GenerateToken(user.ID, a.jwtSecret, a.tokenTTL)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (a *authImpl) VerifyToken(ctx context.Context, token string) (uint, error) {
	claims, err := auth.ParseToken(a.jwtSecret, token)
	if err != nil {
		return 0, err
	}
	return claims.UserID, nil
}
