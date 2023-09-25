package request

import (
	"strings"

	"github.com/digicon-trap1-2023/backend/domain"
)

type GetDocumentsRequest struct {
	Tags string `json:"tags" query:"tags"`
}

type GetDocumentsResponse struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	File       string `json:"file"`
	BookMarked bool   `json:"bookMarked"`
	Referenced bool   `json:"referenced"`
}

func (r *GetDocumentsRequest) ParseTags() []string {
	tags := strings.Split(r.Tags, ",")
	return tags
}

func DocumentToGetDocumentsResponse(document *domain.Document) GetDocumentsResponse {
	return GetDocumentsResponse{
		Id:         document.Id.String(),
		Title:      document.Title,
		File:       document.File,
		BookMarked: document.BookMarked,
		Referenced: document.Referenced,
	}
}

func DocumentsToGetDocumentsResponse(documents []*domain.Document) []GetDocumentsResponse {
	getDocumentsResponse := make([]GetDocumentsResponse, len(documents))

	for i, document := range documents {
		getDocumentsResponse[i] = DocumentToGetDocumentsResponse(document)
	}

	return getDocumentsResponse
}

type GetDocumentRequest struct {
	Id string `json:"id" query:"id"`
}

type GetDocumentResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	File        string `json:"file"`
	Tags        []Tag  `json:"tags"`
	Description string `json:"description"`
	BookMarked  bool   `json:"bookMarked"`
	Referenced  bool   `json:"referenced"`
}

func DocumentToGetDocumentResponse(document *domain.Document) GetDocumentResponse {
	return GetDocumentResponse{
		Id:          document.Id.String(),
		Title:       document.Title,
		File:        document.File,
		Tags:        ConvertTags(document.Tags),
		Description: document.Description,
		BookMarked:  document.BookMarked,
		Referenced:  document.Referenced,
	}
}

