package auth

import "context"

type AuthUC interface {
	Authenticate(ctx context.Context, email string, password string) (string, error)
	VerifyToken(ctx context.Context, token string) error
}
