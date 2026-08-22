package task

import (
	"context"
	"tracker/internal/task/model"
)

type Repository interface {
	CreateTask(ctx context.Context, task *model.Task) (uint, error)
	GetTasks(ctx context.Context, filter model.TaskFilter, page model.Page) (*model.PageResult, error)
}
