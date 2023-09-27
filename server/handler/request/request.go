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
	tags := make([]string, len(request.Tags))
	for i, tag := range request.Tags {
		tags[i] = tag.String()
	}

	return &GetRequestsResponse{
		Id:          request.Id.String(),
		Title:       request.Title,
		Description: request.Description,
		Tags:        tags,
		CreatedBy:   request.CreatedBy.String(),
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
	tags := make([]string, len(request.Tags))
	for i, tag := range request.Tags {
		tags[i] = tag.String()
	}

	return &PostRequestResponse{
		Id:          request.Id.String(),
		Title:       request.Title,
		Description: request.Description,
		Tags:        tags,
		CreatedBy:   request.CreatedBy.String(),
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
	Id    string `json:"id"`
	Title string `json:"title"`
	FileUrl  string `json:"file_url"`
}
