package domain

import "github.com/google/uuid"

type Document struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	FileId      uuid.UUID `json:"fileId"`
}
