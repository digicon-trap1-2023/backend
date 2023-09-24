package handler

import (
	"net/http"

	"github.com/digicon-trap1-2023/backend/handler/request"
	"github.com/digicon-trap1-2023/backend/usecases/service"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type DocumentHandler struct {
	s *service.DocumentService
}

func NewDocumentHandler(s *service.DocumentService) *DocumentHandler {
	return &DocumentHandler{s}
}

func (h *DocumentHandler) GetDocuments(c echo.Context) error {
	var req request.GetDocumentsRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userIdString := c.Get("userId").(string)
	userId, err := uuid.Parse(userIdString)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	documents, err := h.s.GetDocuments(userId, req.ParseTags())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, documents)
}
