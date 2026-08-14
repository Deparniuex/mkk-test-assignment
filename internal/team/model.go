package team

import "time"

type TeamModel struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	CreatedBy int64     `db:"created_by"`
	CreatedAt time.Time `db:"created_at"`
}

type TeamMemberModel struct {
	// unique pair [teamID | userID]
	TeamID int64 `db:"team_id"`
	UserID int64 `db:"user_id"`
	Role   Role  `db:"role"`
}

type Role int

const (
	Owner Role = iota
	Admin
	Member
)
