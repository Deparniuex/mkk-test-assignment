package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"tracker/internal/base/database"
	"tracker/internal/user"
)

type repository struct {
	mysql *sql.DB
}

func NewUserRepository(db *sql.DB) user.Repository {
	return &repository{mysql: db}
}

var usersTable = `users`

func (r *repository) CreateUser(ctx context.Context, user *user.UserModel) error {
	query := fmt.Sprintf(`
	INSERT INTO %s (
		email,
	    password_hash,
	    name,
		created_at
	)
	VALUES (?, ?, ?, ?)
    `, usersTable)
	_, err := r.mysql.ExecContext(ctx, query, user.Email, user.PasswordHash, user.Name, user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUserByID(ctx context.Context, id uint) (*user.UserModel, error) {
	query := fmt.Sprintf(`
		SELECT * FROM %s WHERE id = ?
	`, usersTable)
	user := &user.UserModel{}
	err := r.mysql.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, database.ErrNotFound
	case err != nil:
		return nil, err
	default:
		return user, nil
	}
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*user.UserModel, error) {
	query := fmt.Sprintf(`
		SELECT * FROM %s WHERE email = ?
	`, usersTable)
	user := &user.UserModel{}
	err := r.mysql.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.Name, &user.CreatedAt)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, database.ErrNotFound
	case err != nil:
		return nil, err
	default:
		return user, nil
	}
}
