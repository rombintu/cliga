package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Server) pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "PONG!")
}

// func (s *Server) sprintsHandler(c echo.Context) error {
// 	id := c.Param("id")
// 	if id == "" {
// 		return c.String(http.StatusBadRequest, "sprint id is empty")
// 	}

// 	idint, err := strconv.Atoi(id)
// 	if err != nil {
// 		return c.String(http.StatusBadRequest, "sprint id must be integer")
// 	}

// 	sprint, err := s.store.FetchSprint(idint)
// 	if err != nil {
// 		return c.String(http.StatusInternalServerError, "server error")
// 	}
// 	return c.JSON(http.StatusOK, sprint)
// }
