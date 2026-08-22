package impl

import (
	"context"
	"tracker/internal/task"
	"tracker/internal/task/model"
	"tracker/internal/team"
	"tracker/internal/user"
)

type usecase struct {
	teamRepo  team.Repository
	taskCache task.Cache
	userRepo  user.Repository
	taskRepo  task.Repository
}

func NewTaskUC(teamRepo team.Repository, userRepo user.Repository, taskRepo task.Repository) task.UseCase {
	return &usecase{
		teamRepo: teamRepo,
		userRepo: userRepo,
		taskRepo: taskRepo,
	}
}

func (u *usecase) CreateTask(ctx context.Context, task *model.Task) (uint, error) {
	if task.Status == "" {
		task.Status = model.StatusNew
	}
	_, err := u.teamRepo.GetMember(ctx, task.AssigneeID, task.TeamID)
	if err != nil {
		return 0, err
	}
	return u.taskRepo.CreateTask(ctx, task)
}

func (u *usecase) GetTasks(ctx context.Context, userID uint, filter model.TaskFilter, page model.Page) (*model.PageResult, error) {
	_, err := u.teamRepo.GetMember(ctx, userID, filter.TeamID)
	if err != nil {
		return nil, err
	}
	return u.taskRepo.GetTasks(ctx, filter, page)
}
