package handler

import (
	"context"

	"github.com/digicon-trap1-2023/backend/util"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		role := c.Request().Header.Get("Role")
		ctx = context.WithValue(ctx, util.UserKey, util.SampleUserID().String())
		newCtx := context.WithValue(ctx, util.RoleKey, role)
		c.SetRequest(c.Request().WithContext(newCtx))
		return next(c)
	}
}
