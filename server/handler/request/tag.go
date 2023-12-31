package request

import "github.com/digicon-trap1-2023/backend/domain"

type Tag struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func ConvertTag(tag *domain.Tag) Tag {
	return Tag{
		Id:   tag.Id.String(),
		Name: tag.Name,
	}
}

type GetTagsResponse []Tag

type PostTagRequest struct {
	Name string `json:"name"`
}

func ConvertTags(tags []*domain.Tag) []Tag {
	getTagsResponse := make([]Tag, len(tags))

	for i, tag := range tags {
		getTagsResponse[i] = ConvertTag(tag)
	}

	return getTagsResponse
}
