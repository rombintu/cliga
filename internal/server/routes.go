package server

import (
	"net/http"
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/rombintu/checker-sprints/internal/storage"
)

func (s *Server) pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "PONG!")
}

func (s *Server) userUpdateHandler(c echo.Context) error {
	var user storage.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	if user.Login == "" || user.Anchor == "" {
		return c.String(http.StatusBadRequest, "login or anchor are empty")
	}

	oldUser, err := s.store.UserFetch(user.Login)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if oldUser.Anchor != user.Anchor {
		return c.String(http.StatusForbidden, "anchors are different")
	}

	if reflect.DeepEqual(user.Sprints, oldUser.Sprints) {
		return c.String(http.StatusOK, "no changes")
	} else {
		user.Sprints = oldUser.Sprints
	}

	if err := s.store.UserUpsert(user); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "updated")
}
