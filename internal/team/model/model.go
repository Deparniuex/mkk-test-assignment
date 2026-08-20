package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Team struct {
	ID        uint      `db:"id"`
	Name      string    `db:"name"`
	CreatedBy uint      `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
}

type TeamMember struct {
	// unique pair [teamID | userID]
	TeamID uint `db:"team_id"`
	UserID uint `db:"user_id"`
	Role   Role `db:"role"`
}

type Role int

const (
	RoleUnspecified Role = iota
	Owner
	Admin
	Member
)

var roleToString = map[Role]string{
	Owner:  "owner",
	Admin:  "admin",
	Member: "member",
}

var stringToRole = map[string]Role{
	"owner":  Owner,
	"admin":  Admin,
	"member": Member,
}

func (r Role) String() string {
	s, ok := roleToString[r]
	if !ok {
		return "unknown"
	}
	return s
}

func (r Role) Value() (driver.Value, error) {
	s, ok := roleToString[r]
	if !ok {
		return nil, fmt.Errorf("invalid role: %d", r)
	}
	return s, nil
}

func (r *Role) Scan(value interface{}) error {
	var s string
	switch v := value.(type) {
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		return fmt.Errorf("unsupported type for role: %T", value)
	}

	role, ok := stringToRole[s]
	if !ok {
		return fmt.Errorf("unknown role: %s", s)
	}

	*r = role
	return nil
}

func (r Role) CanInvite() bool {
	switch r {
	case Owner, Admin:
		return true
	default:
		return false
	}
}

func (r Role) Assignable() bool {
	return r != Owner
}

func ParseRole(s string) (Role, error) {
	r, ok := stringToRole[s]
	if !ok {
		return 0, fmt.Errorf("unknown role: %s", s)
	}
	return r, nil
}
