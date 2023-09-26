package request

type GetRequestsResponse struct {
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	CreatedBy   string   `json:"created_by"`
}

type PostRequestRequest struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
}

type DeleteRequestRequest struct {
	Id string `json:"id" param:"id"`
}
