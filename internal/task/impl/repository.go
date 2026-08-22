package impl

import (
	"context"
	"database/sql"
	"fmt"
	"tracker/internal/task"
	"tracker/internal/task/model"
)

type repository struct {
	sqlDB *sql.DB
}

func NewTaskRepository(sqlDB *sql.DB) task.Repository {
	return &repository{
		sqlDB: sqlDB,
	}
}

var (
	taskTable        = "tasks"
	taskHistoryTable = "task_history"
)

func (r *repository) CreateTask(ctx context.Context, task *model.Task) (uint, error) {
	queryTask := fmt.Sprintf(`
	INSERT INTO %s (
	    team_id,
	    title,
		description,
	    status,
	    created_by,
	    assignee_id,
	    version
	)
	VALUES (?, ?, ?, ?, ?, ?, ?)
`, taskTable)

	res, execErr := r.sqlDB.ExecContext(ctx, queryTask, task.TeamID,
		task.Title, task.Description, task.Status, task.CreatedBy,
		task.AssigneeID, task.Version)
	if execErr != nil {
		return 0, execErr
	}
	taskID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint(taskID), nil
}

func (r *repository) GetTasks(ctx context.Context, filter model.TaskFilter, page model.Page) (*model.PageResult, error) {
	if page.Limit <= 0 {
		page.Limit = 20
	}

	query := fmt.Sprintf(`
		SELECT t.id, t.title, t.status, t.assignee_id, t.team_id, t.created_at
		FROM %s t
		WHERE t.id > ?
	`, taskTable)
	args := []any{page.Cursor}

	if filter.AssigneeID != 0 {
		query += " AND t.assignee_id = ?"
		args = append(args, filter.AssigneeID)
	}
	if filter.TeamID != 0 {
		query += " AND t.team_id = ?"
		args = append(args, filter.TeamID)
	}
	if filter.Status != nil {
		query += " AND t.status = ?"
		args = append(args, *filter.Status)
	}

	query += " ORDER BY t.id ASC LIMIT ?"
	args = append(args, page.Limit+1)

	rows, err := r.sqlDB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Status, &t.AssigneeID, &t.TeamID, &t.CreatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, &t)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	hasMore := len(tasks) > page.Limit
	if hasMore {
		tasks = tasks[:page.Limit]
	}

	var nextCursor uint
	if len(tasks) > 0 {
		nextCursor = tasks[len(tasks)-1].ID
	}

	return &model.PageResult{
		Tasks:      tasks,
		NextCursor: nextCursor,
		HasMore:    hasMore,
	}, nil
}
