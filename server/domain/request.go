package domain

import "github.com/google/uuid"

type Request struct {
	Id          uuid.UUID   `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Tags        []uuid.UUID `json:"tags"`
	CreatedBy   uuid.UUID   `json:"created_by"`
}
