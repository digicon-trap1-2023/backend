package handler

import (
	"context"
	"net/http"

	"github.com/digicon-trap1-2023/backend/handler/request"
	"github.com/digicon-trap1-2023/backend/usecases/service"
	"github.com/digicon-trap1-2023/backend/util"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	s *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{s}
}

func (h *AuthHandler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		role := c.Request().Header.Get("X-Role")
		// userId := c.Request().Header.Get("X-UserId")
		ctx = context.WithValue(ctx, util.UserKey, util.SampleUserID().String())
		newCtx := context.WithValue(ctx, util.RoleKey, role)
		c.SetRequest(c.Request().WithContext(newCtx))
		return next(c)
	}
}

func (h *AuthHandler) GetMe(c echo.Context) error {
	userId, err := request.GetUserId(c)
	if err != nil {
		return err
	}
	user, err := h.s.GetUser(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, request.UserToGetMeResponse(user))
}
