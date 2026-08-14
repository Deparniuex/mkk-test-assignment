package http

import (
	"tracker/internal/config"

	"github.com/gin-gonic/gin"
)

func newRouter(_ *config.Config, handlers Handlers) *gin.Engine {
	r := gin.Default()
	// ENDPOINTS
	return r
}
