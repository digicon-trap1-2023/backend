package domain

import "github.com/google/uuid"

type BookMark struct {
	Id         uuid.UUID `json:"id"`
	UserId     uuid.UUID `json:"userId"`
	DocumentId uuid.UUID `json:"documentId"`
}
