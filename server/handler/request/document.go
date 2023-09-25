package request

import (
	"strings"

	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/google/uuid"
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

type PostDocumentRequest struct {
	Title       string   `json:"title"`
	TagIds      []string `json:"tags"`
	Description string   `json:"description"`
}

func GetTagIds(tagIdStrings []string) ([]uuid.UUID, error) {
	tagIds := make([]uuid.UUID, len(tagIdStrings))

	for i, tagId := range tagIdStrings {
		id, err := uuid.Parse(tagId)
		if err != nil {
			return nil, err
		}

		tagIds[i] = id
	}

	return tagIds, nil
}

type PatchDocumentRequest struct {
	Id          string   `json:"id" param:"id"`
	Title       string   `json:"title"`
	TagIds      []string `json:"tags"`
	Description string   `json:"description"`
}

func (r *PatchDocumentRequest) GetTagIds() ([]uuid.UUID, error) {
	tagIds := make([]uuid.UUID, len(r.TagIds))

	for i, tagId := range r.TagIds {
		id, err := uuid.Parse(tagId)
		if err != nil {
			return nil, err
		}

		tagIds[i] = id
	}

	return tagIds, nil
}
