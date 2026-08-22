package task

import (
	"errors"
)

var (
	ErrMemberNotFound = errors.New("user is not a part of this team")
)
