package auth

import (
	"errors"
	"net/http"
	"tracker/internal/base/response"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authUC AuthUC
}

func NewAuthHandler(authUC AuthUC) *AuthHandler {
	return &AuthHandler{authUC: authUC}
}

type logInRequest struct {
	Email    string `json:"email" binding:"required" example:"example@gmail.com"`
	Password string `json:"password" binding:"required" example:"qwerty"`
}

func (h *AuthHandler) LogIn(ctx *gin.Context) {
	var req logInRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.WriteResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.authUC.Authenticate(ctx, req.Email, req.Password)
	if err != nil {
		response.WriteResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	tokenStruct := gin.H{
		"token": token,
	}
	ctx.JSON(http.StatusOK, &response.Response{
		Code:    http.StatusOK,
		Message: "user successfully authenticated",
		Body:    tokenStruct,
	})
}

func (h *AuthHandler) VerifyToken(ctx *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ctx.GetHeader("Authorization")
		err := h.authUC.VerifyToken(ctx, token)
		if err != nil {
			switch {
			case errors.Is(err, ErrInvalidToken):
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
					Code:    http.StatusUnauthorized,
					Message: "token is invalid",
				})
				return
			case errors.Is(err, ErrTokenExpired):
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
					Code:    http.StatusUnauthorized,
					Message: "token is expired",
				})
			default:
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, response.Response{
					Code:    http.StatusUnauthorized,
					Message: err.Error(),
				})
				return
			}
		}
		ctx.Next()
	}
}
