package server

import (
	"context"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rombintu/checker-sprints/internal/cli"
	"github.com/rombintu/checker-sprints/internal/config"
	"github.com/rombintu/checker-sprints/internal/storage"
)

type Server struct {
	store  *storage.MongodbDriver
	router *echo.Echo
	config config.ServerConfig
}

func NewServer(conf config.ServerConfig, store *storage.MongodbDriver) *Server {
	return &Server{
		store:  store,
		router: echo.New(),
		config: conf,
	}
}

func (s *Server) Configure() {
	s.ConfigureStore()
	s.configureLogger()
	s.configureRouter()
	s.configureSprints()
}

func (s *Server) ConfigureStore() {
	ctx := context.Background()
	if err := s.store.Open(ctx); err != nil {
		slog.Error(err.Error())
		s.store.Close(ctx)
		os.Exit(1)
	}
}

func (s *Server) configureRouter() {
	s.router.GET("/", s.pingHandler)
	s.router.GET("/sprints/:num", s.sprintsHandler)
	s.router.POST("/users/sprint/:num", s.userSprintHandler)
}

func (s *Server) configureLogger() {
	s.router.Use(middleware.Logger())
}

func (s *Server) configureSprints() {
	cli.SprintsInit()
}

func (s *Server) Start() {
	if err := s.router.Start(":8080"); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
