package impl

import (
	"context"
	"tracker/internal/user"
)

type usecase struct {
	userRepo user.Repository
}

func NewUserUC(userRepo user.Repository) user.UserUC {
	return &usecase{
		userRepo: userRepo,
	}
}

func (u *usecase) CreateUser(ctx context.Context, user *user.UserModel) error {
	if err := u.userRepo.CreateUser(ctx, user); err != nil {
		return err
	}
	return nil
}

func (u *usecase) GetUserByID(ctx context.Context, id uint) (*user.UserModel, error) {
	user, err := u.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usecase) GetUserByEmail(ctx context.Context, email string) (*user.UserModel, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}
