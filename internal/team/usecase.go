package team

import "context"

type TeamUC interface {
	CreateTeam(ctx context.Context, name string, userID int64) error
	GetTeamsByUser(ctx context.Context, userID int64)
	InviteUser(ctx context.Context)
}
