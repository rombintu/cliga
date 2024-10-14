package server

import (
	"errors"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rombintu/checker-sprints/internal/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Server) pingHandler(c echo.Context) error {
	return c.String(http.StatusOK, "PONG!")
}

func (s *Server) userSprintHandler(c echo.Context) error {
	num := c.Param("num")
	if num == "" {
		return c.String(http.StatusBadRequest, "sprint number is empty")
	}
	sprintNum, err := strconv.ParseInt(num, 10, 64)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}
	var user storage.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	if user.Login == "" || user.Anchor == "" {
		return c.String(http.StatusBadRequest, "login or anchor are empty")
	}
	isNewUser := false
	oldUser, err := s.store.UserFetch(user.Login)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			slog.Warn("user not singin. process registration...")
			if err := s.store.UserLogin(user); err != nil {
				return c.String(http.StatusInternalServerError, err.Error())
			}
			isNewUser = true
		} else {
			return c.String(http.StatusInternalServerError, err.Error())
		}
	}

	if !isNewUser && oldUser.Anchor != user.Anchor {
		return c.String(http.StatusForbidden, "handprints are incorrect")
	}

	for _, spr := range oldUser.Sprints {
		if spr.ID == sprintNum {
			return c.String(http.StatusOK, "no change")
		}
	}

	sprintok := storage.Sprint{
		ID:        sprintNum,
		IsDone:    true,
		UpdatedAt: time.Now(),
	}
	if err := s.store.UserPushSprint(user.Login, sprintok); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.String(http.StatusOK, "updated")
}

func (s *Server) userGetHandler(c echo.Context) error {
	username := c.Param("username")
	if username == "" {
		return c.String(http.StatusBadRequest, "username is empty")
	}

	user, err := s.store.UserFetch(username)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
