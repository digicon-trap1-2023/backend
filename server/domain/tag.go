package domain

import "github.com/google/uuid"

type Tag struct {
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type TagDocument struct {
	Id         uuid.UUID `json:"id"`
	TagId      uuid.UUID `json:"tagId"`
	DocumentId uuid.UUID `json:"documentId"`
}
