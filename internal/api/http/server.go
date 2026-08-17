package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Host string `env:"REDIS_HOST" envDefault:"localhost:6379"`
	Port string `env:"PORT" envDefault:"8080"`
}

type Server struct {
	config *Config
	server *http.Server
	notify chan error
}

func NewServer(cfg *Config, handlers Handlers) *Server {
	gin.SetMode(gin.DebugMode)

	router := newRouter(handlers)

	return &Server{
		server: &http.Server{
			Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
			Handler: router.Handler(),
		},
		config: cfg,
		notify: make(chan error, 1),
	}
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*1000)
	defer cancel()

	return s.server.Shutdown(ctx)
}

func (s *Server) Notify() <-chan error {
	return s.notify
}
