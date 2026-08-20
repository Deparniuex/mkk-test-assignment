package impl

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"tracker/internal/team"
	"tracker/internal/team/model"
)

type repository struct {
	db *sql.DB
}

func NewTeamRepository(db *sql.DB) team.Repository {
	return &repository{db: db}
}

var teamTable = "teams"
var teamMemberTable = "team_members"

func (r *repository) CreateTeam(ctx context.Context, tm *model.Team) (uint, error) {
	queryTeam := fmt.Sprintf(`
	INSERT INTO %s (
	                name, 
	                created_by
	) 
	VALUES (?, ?)
	`, teamTable)

	queryMember := fmt.Sprintf(`
	INSERT INTO %s (
	                team_id,
	                user_id,
	                role
	)
	VALUES (?, ?, ?)
	`, teamMemberTable)
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	res, execErr := tx.ExecContext(ctx, queryTeam, tm.Name, tm.CreatedBy)
	if execErr != nil {
		return 0, execErr
	}
	teamID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	_, execErr = tx.ExecContext(ctx, queryMember, teamID, tm.CreatedBy, model.Owner)
	if execErr != nil {
		return 0, execErr
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return uint(teamID), nil
}

func (r *repository) GetTeamsByUser(ctx context.Context, userID uint) ([]*model.Team, error) {
	query := fmt.Sprintf(`
	SELECT t.id, t.name, t.created_by, t.created_at
	FROM %s t
	JOIN %s tm ON tm.team_id = t.id
	WHERE tm.user_id = ?
	`, teamTable, teamMemberTable)
	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var members []*model.Team
	for rows.Next() {
		var member model.Team
		err := rows.Scan(&member.ID, &member.Name, &member.CreatedBy, &member.CreatedAt)
		if err != nil {
			return nil, err
		}
		members = append(members, &member)
	}
	return members, rows.Err()
}

func (r *repository) GetMember(ctx context.Context, userID uint, teamID uint) (*model.TeamMember, error) {
	query := fmt.Sprintf(`
		SELECT team_id, user_id, role
		FROM %s
		WHERE team_id = ? AND user_id = ?
	`, teamMemberTable)

	row := r.db.QueryRowContext(ctx, query, teamID, userID)
	member := &model.TeamMember{}
	err := row.Scan(&member.TeamID, &member.UserID, &member.Role)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, team.ErrMemberNotFound
		default:
			return nil, err
		}
	}
	return member, err
}

func (r *repository) InviteUser(ctx context.Context, userID uint, teamID uint, newMember uint, role model.Role) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (team_id, user_id, role)
		VALUES (?, ?, ?)
	`, teamMemberTable)
	_, err := r.db.ExecContext(ctx, query, teamID, newMember, role)
	if err != nil {
		return err
	}
	return nil
}
