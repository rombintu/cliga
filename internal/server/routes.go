package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) rootHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
