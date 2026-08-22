package model

import "time"

type Task struct {
	ID          uint      `db:"id"`
	TeamID      uint      `db:"team_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	Status      Status    `db:"status"`
	CreatedBy   uint      `db:"created_by"`
	AssigneeID  uint      `db:"assignee_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	ClosedAt    time.Time `db:"closed_at"`
	Version     uint      `db:"version"`
}

type Status string

const (
	StatusNew        Status = "new"
	StatusTodo       Status = "todo"
	StatusInProgress Status = "in_progress"
	StatusResolved   Status = "resolved"
	StatusClosed     Status = "closed"
)
