package handler

import (
	"net/http"

	"github.com/digicon-trap1-2023/backend/handler/request"
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

	return c.JSON(http.StatusOK, request.RequestsToGetRequestsResponse(requests))
}

func (h *RequestHandler) PostRequest(c echo.Context) error {
	var req request.PostRequestRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId, err := request.GetUserId(c)
	if err != nil {
		return err
	}

	tagIds, err := request.GetTagIds(req.Tags)
	if err != nil {
		return err
	}

	domainRequest, err := h.s.CreateRequest(userId, tagIds, req.Title, req.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, request.RequestToPostRequestResponse(domainRequest))
}
