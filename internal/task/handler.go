package task

import (
	"errors"
	"net/http"
	"strconv"
	"tracker/internal/base/response"
	"tracker/internal/task/model"
	"tracker/internal/team"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	taskUC UseCase
}

func NewTaskHandler(taskUC UseCase) *TaskHandler {
	return &TaskHandler{taskUC: taskUC}
}

type CreateTaskRequest struct {
	TeamID      uint   `json:"team_id" binding:"required" example:"1"`
	Title       string `json:"title" binding:"required" example:"New task"`
	Description string `json:"description" example:"New task"`
	Status      string `json:"status" example:"New task"`
	AssigneeID  uint   `json:"assignee_id" binding:"required" example:"1"`
}

func (h *TaskHandler) CreateTask(ctx *gin.Context) {
	var req CreateTaskRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.WriteResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	createdBy := ctx.GetUint("userID")
	task := model.Task{
		TeamID:     req.TeamID,
		Title:      req.Title,
		Status:     model.Status(req.Status),
		CreatedBy:  createdBy,
		AssigneeID: req.AssigneeID,
	}
	taskID, err := h.taskUC.CreateTask(ctx, &task)
	if err != nil {
		switch {
		case errors.Is(err, team.ErrMemberNotFound):
			response.WriteResponse(ctx, http.StatusForbidden, ErrMemberNotFound.Error())
			return
		default:
			response.WriteResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ctx.JSON(http.StatusCreated, response.Response{
		Code:    http.StatusCreated,
		Message: "task successfully created",
		Body:    gin.H{"task_id": taskID},
	})
}

func (h *TaskHandler) GetTasks(ctx *gin.Context) {
	cursor, err := model.DecodeCursor(ctx.Query("cursor"))
	if err != nil {
		response.WriteResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "20"))

	var filter model.TaskFilter
	if v := ctx.Query("assignee_id"); v != "" {
		id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			response.WriteResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		filter.AssigneeID = uint(id)
	}

	if v := ctx.Query("team_id"); v != "" {
		id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			response.WriteResponse(ctx, http.StatusBadRequest, err.Error())
			return
		}
		filter.TeamID = uint(id)
	}
	if v := ctx.Query("status"); v != "" {
		status := model.Status(v)
		filter.Status = &status
	}

	userID := ctx.GetUint("userID")

	result, err := h.taskUC.GetTasks(ctx, userID, filter, model.Page{
		Cursor: cursor,
		Limit:  limit,
	})
	if err != nil {
		switch {
		case errors.Is(err, team.ErrMemberNotFound):
			response.WriteResponse(ctx, http.StatusForbidden, ErrMemberNotFound.Error())
			return
		default:
			response.WriteResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ctx.JSON(http.StatusOK, response.Response{
		Code:    http.StatusOK,
		Message: "tasks successfully retrieved",
		Body: gin.H{
			"tasks":       result.Tasks,
			"next_cursor": model.EncodeCursor(result.NextCursor),
			"has_more":    result.HasMore,
		},
	})
}
