package request

type Tag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetTagsResponse []Tag

type CreateTagRequest struct {
	Name string `json:"name"`
}

type CreateTagResponse Tag
