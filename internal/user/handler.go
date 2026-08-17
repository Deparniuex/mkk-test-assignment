package user

import (
	"net/http"
	"time"
	"tracker/internal/base/response"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUC UserUC
}

func NewUserHandler(userUC UserUC) *UserHandler {
	return &UserHandler{
		userUC: userUC,
	}
}

type createUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=50" example:"John Doe"`
	Email    string `json:"email"    binding:"required,min=2,max=50" example:"example@gmail.com"`
	Password string `json:"password" binding:"required,min=6,max=32" example:"qwerty"`
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var req createUserRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		response.WriteResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := UserModel{
		Name:      req.Name,
		Email:     req.Email,
		CreatedAt: time.Now(),
	}

	err = user.SetPassword(req.Password)
	if err != nil {
		response.WriteResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.userUC.CreateUser(ctx, &user)
	if err != nil {
		response.WriteResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, &response.Response{
		Code:    http.StatusCreated,
		Message: "user successfully created",
	})
}
