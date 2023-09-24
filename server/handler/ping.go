package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type PingHandler struct {
	conn *gorm.DB
}

func NewPingHandler(db *gorm.DB) *PingHandler {
	return &PingHandler{db}
}

func (h *PingHandler) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}
