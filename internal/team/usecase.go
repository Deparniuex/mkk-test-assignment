package team

import (
	"context"
	"tracker/internal/team/model"
)

type UC interface {
	CreateTeam(ctx context.Context, team *model.Team) (uint, error)
	GetTeamsByUser(ctx context.Context, userID uint) ([]*model.Team, error)
	InviteUser(ctx context.Context, userID uint, teamID uint, newMember uint, role model.Role) error
}
