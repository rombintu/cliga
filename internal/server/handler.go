package server

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rombintu/checker-sprints/internal/config"
)

type Server struct {
	// store  *storage.Storage
	router *echo.Echo
	config config.ServerConfig
}

func NewServer(conf config.ServerConfig) *Server {
	return &Server{
		router: echo.New(),
		config: conf,
	}
}

func (s *Server) Configure() {
	s.configureLogger()
	s.configureRouter()
}

func (s *Server) configureRouter() {
	s.router.GET("/", s.rootHandler)
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
