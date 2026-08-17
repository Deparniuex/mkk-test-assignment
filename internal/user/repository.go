package user

import "context"

type Repository interface {
	CreateUser(ctx context.Context, user *UserModel) error
	GetUserByID(ctx context.Context, id uint) (*UserModel, error)
	GetUserByEmail(ctx context.Context, email string) (*UserModel, error)
}
