package domain

import "github.com/google/uuid"

type Document struct {
	Id             uuid.UUID `json:"id"`
	UserName       string    `json:"user_name"`
	Title          string    `json:"title"`
	File           string    `json:"file"`
	FileHeight     int       `json:"file_height"`
	FileWidth      int       `json:"file_width"`
	Description    string    `json:"description"`
	Tags           []*Tag    `json:"tags"`
	BookMarked     bool      `json:"bookmarked"`
	Referenced     bool      `json:"referenced"`
	ReferenceUsers []string  `json:"reference_users"`
}

type Size struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
