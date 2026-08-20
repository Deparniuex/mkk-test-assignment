package team

import "errors"

var (
	ErrForbidden      = errors.New("not enough permissions")
	ErrAlreadyMember  = errors.New("already member of this team")
	ErrMemberNotFound = errors.New("member not found")
	ErrInvalidRole    = errors.New("invalid role")
)
