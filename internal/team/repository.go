package team

import (
	"context"
	"tracker/internal/team/model"
)

type Repository interface {
	CreateTeam(ctx context.Context, team *model.Team) error
	GetMember(ctx context.Context, userID uint, teamID uint) (*model.TeamMember, error)
	GetTeamsByUser(ctx context.Context, userID uint) ([]*model.Team, error)
	InviteUser(ctx context.Context, userID uint, teamID uint, newMember uint, role model.Role) error
}
