package user

import "context"

type UserUC interface {
	CreateUser(ctx context.Context, user *UserModel) error
	GetUser(ctx context.Context, id string) (*UserModel, error)
}
