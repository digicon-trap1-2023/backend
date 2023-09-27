package request

import "github.com/digicon-trap1-2023/backend/domain"

type GetRequestsResponse struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	CreatedBy   string   `json:"created_by"`
}

func RequestsToGetRequestsResponse(requests []*domain.Request) []*GetRequestsResponse {
	getRequestsResponse := make([]*GetRequestsResponse, len(requests))

	for i, request := range requests {
		getRequestsResponse[i] = RequestToGetRequestsResponse(request)
	}

	return getRequestsResponse
}

func RequestToGetRequestsResponse(request *domain.Request) *GetRequestsResponse {
	return &GetRequestsResponse{
		Id:          request.Id.String(),
		Title:       request.Title,
		Description: request.Description,
		Tags:        request.TagNames,
		CreatedBy:   request.CreatedUserName,
	}
}

type PostRequestRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type PostRequestResponse struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	CreatedBy   string   `json:"created_by"`
}

func RequestToPostRequestResponse(request *domain.Request) *PostRequestResponse {
	return &PostRequestResponse{
		Id:          request.Id.String(),
		Title:       request.Title,
		Description: request.Description,
		Tags:        request.TagNames,
		CreatedBy:   request.CreatedUserName,
	}
}

type DeleteRequestRequest struct {
	Id string `json:"id" param:"id"`
}

type GetRequestsWithDocumentResponse struct {
	Id        string              `json:"id"`
	Title     string              `json:"title"`
	Documents []*DocumentResponse `json:"documents"`
}

type DocumentResponse struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	FileUrl    string `json:"file_url"`
	FileHeight int    `json:"file_height"`
	FileWidth  int    `json:"file_width"`
}

func RequestsToGetRequestsWithDocumentResponse(request []*domain.Request) []*GetRequestsWithDocumentResponse {
	getRequestsWithDocumentResponse := make([]*GetRequestsWithDocumentResponse, len(request))

	for i, request := range request {
		getRequestsWithDocumentResponse[i] = RequestToGetRequestsWithDocumentResponse(request)
	}

	return getRequestsWithDocumentResponse
}

func RequestToGetRequestsWithDocumentResponse(request *domain.Request) *GetRequestsWithDocumentResponse {
	documents := make([]*DocumentResponse, len(request.RelatedDocuments))
	for i, document := range request.RelatedDocuments {
		documents[i] = DocumentToDocumentResponse(document)
	}

	return &GetRequestsWithDocumentResponse{
		Id:        request.Id.String(),
		Title:     request.Title,
		Documents: documents,
	}
}

func DocumentToDocumentResponse(document *domain.Document) *DocumentResponse {
	return &DocumentResponse{
		Id:         document.Id.String(),
		Title:      document.Title,
		FileUrl:    document.File,
		FileHeight: document.FileHeight,
		FileWidth:  document.FileWidth,
	}
}
