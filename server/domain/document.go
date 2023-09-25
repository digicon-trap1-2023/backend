package domain

import "github.com/google/uuid"

type Document struct {
	Id          uuid.UUID `json:"id"`
	UserId      uuid.UUID `json:"userId"`
	Title       string    `json:"title"`
	File        string    `json:"file"`
	Description string    `json:"description"`
	Tags        []*Tag    `json:"tags"`
	BookMarked  bool      `json:"bookmarked"`
	Referenced  bool      `json:"referenced"`
}
