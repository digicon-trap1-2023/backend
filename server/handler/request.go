package handler

import (
	"net/http"

	"github.com/digicon-trap1-2023/backend/usecases/service"
	"github.com/labstack/echo/v4"
)

type RequestHandler struct {
	s *service.RequestService
}

func NewRequestHandler(s *service.RequestService) *RequestHandler {
	return &RequestHandler{s}
}

func (h *RequestHandler) GetRequests(c echo.Context) error {
	requests, err := h.s.GetRequests()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, requests)
}
