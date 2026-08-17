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
	}
	return r
}
