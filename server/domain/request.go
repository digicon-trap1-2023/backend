package domain

import "github.com/google/uuid"

type Request struct {
	Id               uuid.UUID   `json:"id"`
	Title            string      `json:"title"`
	Description      string      `json:"description"`
	Tags             []uuid.UUID `json:"tags"`
	TagNames		 []string    `json:"tag_names"`
	CreatedBy        uuid.UUID   `json:"created_by"`
	RelatedDocuments []*Document `json:"related_document"`
}

func NewRequest(title string, description string, tags []uuid.UUID, createdBy uuid.UUID) *Request {
	return &Request{
		Id:          uuid.New(),
		Title:       title,
		Description: description,
		Tags:        tags,
		CreatedBy:   createdBy,
	}
}
