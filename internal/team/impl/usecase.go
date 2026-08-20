package impl

import (
	"context"
	"errors"
	"tracker/internal/team"
	"tracker/internal/team/model"
	"tracker/internal/user"
)

type usecase struct {
	teamRepo       team.Repository
	userRepository user.Repository
}

func NewTeamUC(teamRepo team.Repository, userRepo user.Repository) team.UC {
	return &usecase{
		teamRepo:       teamRepo,
		userRepository: userRepo,
	}
}

func (t *usecase) CreateTeam(ctx context.Context, team *model.Team) error {
	if err := t.teamRepo.CreateTeam(ctx, team); err != nil {
		return err
	}
	return nil
}

func (t *usecase) GetTeamsByUser(ctx context.Context, userID uint) ([]*model.Team, error) {
	teams, err := t.teamRepo.GetTeamsByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (t *usecase) InviteUser(ctx context.Context, userID uint, teamID uint, newMember uint, role model.Role) error {
	inviter, err := t.teamRepo.GetMember(ctx, userID, teamID)
	if err != nil {
		if errors.Is(err, team.ErrMemberNotFound) {
			return team.ErrForbidden
		}
		return err
	}
	if ok := inviter.Role.CanInvite(); !ok {
		return team.ErrForbidden
	}
	if role == 0 {
		role = model.Member
	}
	if !role.Assignable() {
		return team.ErrInvalidRole
	}
	_, err = t.userRepository.GetUserByID(ctx, newMember)
	if err != nil {
		return err
	}
	if _, err := t.teamRepo.GetMember(ctx, newMember, teamID); err == nil {
		return team.ErrAlreadyMember
	} else if !errors.Is(err, team.ErrMemberNotFound) {
		return err
	}
	return t.teamRepo.InviteUser(ctx, userID, teamID, newMember, model.Member)
}
