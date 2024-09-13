package server

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rombintu/checker-sprints/internal/config"
	"github.com/rombintu/checker-sprints/internal/storage"
)

type Server struct {
	store  storage.Storage
	router *echo.Echo
	config config.ServerConfig
}

func NewServer(conf config.ServerConfig, store storage.Storage) *Server {
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
}

func (s *Server) ConfigureStore() {
	if err := s.store.Open(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func (s *Server) configureRouter() {
	s.router.GET("/", s.pingHandler)
	// s.router.GET("/sprints/:id", s.sprintsHandler)
}

func (s *Server) configureLogger() {
	s.router.Use(middleware.Logger())
}

func (s *Server) Start() {
	if err := s.router.Start(":8080"); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
