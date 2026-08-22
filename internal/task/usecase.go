package task

import (
	"context"
	"tracker/internal/task/model"
)

type UseCase interface {
	CreateTask(ctx context.Context, task *model.Task) (uint, error)
	GetTasks(ctx context.Context, userID uint, filter model.TaskFilter, page model.Page) (*model.PageResult, error)
}
