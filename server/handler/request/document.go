package request

import (
	"strings"

	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/google/uuid"
)

type GetDocumentsRequest struct {
	Tags string `json:"tags" query:"tags"`
	Type string `json:"type" query:"type"`
}

type GetDocumentsResponse struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	File       string `json:"file"`
	FileHeight int    `json:"file_height"`
	FileWidth  int    `json:"file_width"`
	BookMarked bool   `json:"bookmarked"`
	Referenced bool   `json:"referenced"`
	UserName   string `json:"user_name"`
}

func (r *GetDocumentsRequest) ParseTags() []string {
	if r.Tags == "" {
		return []string{}
	}
	tags := strings.Split(r.Tags, ",")
	return tags
}

func DocumentToGetDocumentsResponse(document *domain.Document) GetDocumentsResponse {
	return GetDocumentsResponse{
		Id:         document.Id.String(),
		Title:      document.Title,
		File:       document.File,
		FileHeight: document.FileHeight,
		FileWidth:  document.FileWidth,
		BookMarked: document.BookMarked,
		Referenced: document.Referenced,
		UserName:   document.UserName,
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
	Id string `json:"id" param:"id"`
}

type GetDocumentResponse struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	File        string `json:"file"`
	FileHeight  int    `json:"file_height"`
	FileWidth   int    `json:"file_width"`
	Tags        []Tag  `json:"tags"`
	Description string `json:"description"`
	BookMarked  bool   `json:"bookmarked"`
	Referenced  bool   `json:"referenced"`
	UserName    string `json:"user_name"`
}

func DocumentToGetDocumentResponse(document *domain.Document) GetDocumentResponse {
	return GetDocumentResponse{
		Id:          document.Id.String(),
		Title:       document.Title,
		File:        document.File,
		FileHeight:  document.FileHeight,
		FileWidth:   document.FileWidth,
		Tags:        ConvertTags(document.Tags),
		Description: document.Description,
		BookMarked:  document.BookMarked,
		Referenced:  document.Referenced,
		UserName:    document.UserName,
	}
}

func DocumentsToGetOtherDocumentsResponse(documents []*domain.Document) []GetOtherDocumentsResponse {
	getOtherDocumentsResponse := make([]GetOtherDocumentsResponse, len(documents))

	for i, document := range documents {
		getOtherDocumentsResponse[i] = DocumentToGetOtherDocumentsResponse(document)
	}

	return getOtherDocumentsResponse
}

type GetOtherDocumentsResponse struct {
	Id             string   `json:"id"`
	Title          string   `json:"title"`
	File           string   `json:"file"`
	FileHeight     int      `json:"file_height"`
	FileWidth      int      `json:"file_width"`
	Tags           []Tag    `json:"tags"`
	Referenced     bool     `json:"referenced"`
	ReferenceUsers []string `json:"reference_users"`
	UserName       string   `json:"user_name"`
}

func DocumentToGetOtherDocumentsResponse(document *domain.Document) GetOtherDocumentsResponse {
	return GetOtherDocumentsResponse{
		Id:             document.Id.String(),
		Title:          document.Title,
		File:           document.File,
		FileHeight:     document.FileHeight,
		FileWidth:      document.FileWidth,
		Tags:           ConvertTags(document.Tags),
		Referenced:     document.Referenced,
		ReferenceUsers: document.ReferenceUsers,
		UserName:       document.UserName,
	}
}

type PostDocumentRequest struct {
	Title          string   `json:"title"`
	TagIds         []string `json:"tags"`
	Description    string   `json:"description"`
	RelatedRequest string   `json:"related_request; omitempty"`
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
