package team

import (
	"errors"
	"net/http"
	"strconv"
	"tracker/internal/base/database"
	"tracker/internal/base/response"
	"tracker/internal/team/model"

	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	TeamUC UC
}

func NewTeamHandler(teamUC UC) *TeamHandler {
	return &TeamHandler{TeamUC: teamUC}
}

type CreateTeamRequest struct {
	Name string `json:"name" binding:"required" example:"Payment"`
}

func (h *TeamHandler) CreateTeam(ctx *gin.Context) {
	var req CreateTeamRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.WriteResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	userID := ctx.GetUint("userID")
	team := model.Team{Name: req.Name, CreatedBy: userID}
	err := h.TeamUC.CreateTeam(ctx, &team)
	if err != nil {
		response.WriteResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &response.Response{
		Code:    http.StatusCreated,
		Message: "team successfully created",
	})
}

func (h *TeamHandler) GetTeams(ctx *gin.Context) {
	userID := ctx.GetUint("userID")
	teams, err := h.TeamUC.GetTeamsByUser(ctx, userID)
	if err != nil {
		response.WriteResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, &response.Response{
		Code:    http.StatusOK,
		Message: "teams successfully retrieved",
		Body:    teams,
	})
}

type InviteUserRequest struct {
	UserID uint   `json:"user_id" binding:"required" example:"1"`
	Role   string `json:"role" binding:"omitempty,oneof=admin member"`
}

func (h *TeamHandler) InviteUser(ctx *gin.Context) {
	var req InviteUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.WriteResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	userID := ctx.GetUint("userID")
	var role model.Role
	if req.Role != "" {
		var err error
		role, err = model.ParseRole(req.Role)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, &response.Response{
				Code:    http.StatusBadRequest,
				Message: ErrInvalidRole.Error(),
			})
			return
		}
	}

	teamID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &response.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}
	err = h.TeamUC.InviteUser(ctx, userID, uint(teamID), req.UserID, role)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrNotFound):
			response.WriteResponse(ctx, http.StatusNotFound, ErrMemberNotFound.Error())
			return
		case errors.Is(err, ErrAlreadyMember):
			response.WriteResponse(ctx, http.StatusForbidden, ErrAlreadyMember.Error())
			return
		case errors.Is(err, ErrForbidden):
			response.WriteResponse(ctx, http.StatusForbidden, ErrForbidden.Error())
			return
		case errors.Is(err, ErrInvalidRole):
			response.WriteResponse(ctx, http.StatusBadRequest, ErrInvalidRole.Error())
			return
		default:
			response.WriteResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}
	}
	ctx.JSON(http.StatusCreated, &response.Response{
		Code:    http.StatusCreated,
		Message: "member successfully invited",
	})
}
