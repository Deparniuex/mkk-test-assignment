package http

import (
	"github.com/gin-gonic/gin"
)

func newRouter(h Handlers) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", h.User.CreateUser)
		v1.POST("/login", h.Auth.LogIn)

		v1.POST("/teams", h.Auth.VerifyToken(), h.Team.CreateTeam)
		v1.GET("/teams", h.Auth.VerifyToken(), h.Team.GetTeams)
		v1.POST("/teams/:id/invite", h.Auth.VerifyToken(), h.Team.InviteUser)

		v1.POST("/tasks", h.Auth.VerifyToken(), h.Task.CreateTask)
		v1.GET("/tasks", h.Auth.VerifyToken(), h.Task.GetTasks)
	}
	return r
}
