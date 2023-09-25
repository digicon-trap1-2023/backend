package domain

import "github.com/google/uuid"

type Document struct {
	Id          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	File        string    `json:"file"`
	Description string    `json:"description"`
	Tags        []*Tag    `json:"tags"`
	BookMarked  bool      `json:"bookMarked"`
	Referenced  bool      `json:"referenced"`
}
