package task

import "time"

type TaskModel struct {
	ID          int64     `db:"id"`
	TeamID      int64     `db:"team_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      string    `db:"status"`
	CreatedBy   int64     `db:"created_by"`
	AssigneeID  int64     `db:"assignee_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	ClosedAt    time.Time `db:"closed_at"`
}

type TaskHistoryModel struct {
	ID        int64 `db:"id"`
	TaskID    int64 `db:"task_id"`
	ChangedBy int64 `db:"changed_by"`
	Changes
}

type Changes struct {
	Old string `json:"old"`
	New string `json:"new"`
}

type TaskCommentsModel struct {
	ID        int64     `db:"id"`
	TaskID    int64     `db:"task_id"`
	UserID    int64     `db:"user_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
