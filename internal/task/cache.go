package task

import (
	"context"
	"tracker/internal/task/model"
)

type Cache interface {
	SetTask(ctx context.Context, task *model.Task) error
	GetTask(ctx context.Context, id uint) (*model.Task, error)
}
