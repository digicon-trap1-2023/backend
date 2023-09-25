package handler

import (
	"encoding/json"
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

	userId, err := request.GetUserId(c)
	if err != nil {
		return err
	}

	documents, err := h.s.GetDocuments(userId, req.ParseTags())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, request.DocumentsToGetDocumentsResponse(documents))
}

func (h *DocumentHandler) GetDocument(c echo.Context) error {
	var req request.GetDocumentRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	documentId, err := uuid.Parse(req.Id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId, err := request.GetUserId(c)
	if err != nil {
		return err
	}

	document, err := h.s.GetDocument(userId, documentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, request.DocumentToGetDocumentResponse(document))
}

func (h *DocumentHandler) PostDocument(c echo.Context) error {
	title := c.FormValue("title")
	description := c.FormValue("description")
	tagsRaw := c.FormValue("tags")

	var tagIdStrings []string
	err := json.Unmarshal([]byte(tagsRaw), &tagIdStrings)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse tags")
	}

	userId, err := request.GetUserId(c)
	if err != nil {
		return err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tagIds, err := request.GetTagIds(tagIdStrings)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	document, err := h.s.CreateDocument(userId, title, description, tagIds, file)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, document)
}

func (h *DocumentHandler) PatchDocument(c echo.Context) error {
	id := c.Param("id")
	title := c.FormValue("title")
	description := c.FormValue("description")
	tagsRaw := c.FormValue("tags")

	var tagIdStrings []string
	err := json.Unmarshal([]byte(tagsRaw), &tagIdStrings)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Failed to parse tags")
	}
	documentId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId, err := request.GetUserId(c)
	if err != nil {
		return err
	}

	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	tagIds, err := request.GetTagIds(tagIdStrings)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	document, err := h.s.UpdateDocument(userId, documentId, title, description, tagIds, file)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, document)
}

func (h *DocumentHandler) DeleteDocument(c echo.Context) error {
	id := c.Param("id")
	documentId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId, err := request.GetUserId(c)
	if err != nil {
		return err
	}

	err = h.s.DeleteDocument(userId, documentId)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *DocumentHandler) PostBookmark(c echo.Context) error {
	id := c.Param("id")
	documentId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId, err := request.GetUserId(c)
	if err != nil {
		return err
	}

	err = h.s.BookMark(userId, documentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *DocumentHandler) DeleteBookmark(c echo.Context) error {
	id := c.Param("id")
	documentId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId, err := request.GetUserId(c)
	if err != nil {
		return err
	}

	err = h.s.UnBookMark(userId, documentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *DocumentHandler) PostReference(c echo.Context) error {
	id := c.Param("id")
	documentId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId, err := request.GetUserId(c)
	if err != nil {
		return err
	}

	err = h.s.Reference(userId, documentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *DocumentHandler) DeleteReference(c echo.Context) error {
	id := c.Param("id")
	documentId, err := uuid.Parse(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userId, err := request.GetUserId(c)
	if err != nil {
		return err
	}

	err = h.s.UnReference(userId, documentId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
