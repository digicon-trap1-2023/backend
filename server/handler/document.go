package handler

import "github.com/labstack/echo/v4"

type DocumentHandler struct {
}

func NewDocumentHandler() *DocumentHandler {
	return &DocumentHandler{}
}

func (h *DocumentHandler) GetDocuments(c echo.Context) error {
	return nil
}
