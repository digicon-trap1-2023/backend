package handler

import (
	"net/http"

	"github.com/digicon-trap1-2023/backend/handler/request"
	"github.com/digicon-trap1-2023/backend/usecases/service"
	"github.com/labstack/echo/v4"
)

type TagHandler struct {
	s *service.TagService
}

func NewTagHandler(s *service.TagService) *TagHandler {
	return &TagHandler{s}
}

func (h *TagHandler) GetTags(c echo.Context) error {
	tags, err := h.s.GetTags()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	res := make(request.GetTagsResponse, len(tags))
	for i, tag := range tags {
		res[i] = request.ConvertTag(tag)
	}

	return c.JSON(http.StatusOK, res)
}

func (h *TagHandler) PostTag(c echo.Context) error {
	var req request.PostTagRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tag, err := h.s.CreateTag(req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, request.ConvertTag(tag))
}
