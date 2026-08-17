package http

import (
	"tracker/internal/auth"
	"tracker/internal/user"
)

type Handlers struct {
	User *user.UserHandler
	Auth *auth.AuthHandler
}
