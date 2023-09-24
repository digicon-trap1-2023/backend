package model

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/google/uuid"
)

type Document struct {
	Id          string `gorm:"type:char(36);not null;primaryKey"`
	Title       string `gorm:"type:varchar(40)"`
	Description string `gorm:"type:varchar(200)"`
	File        string `gorm:"type:varchar(200)"`
}

func (Document) TableName() string {
	return "documents"
}

func (d *Document) ToDomain(userBookmarks []*BookMark, userReferences []*Reference) (*domain.Document, error) {
	id, err := uuid.Parse(d.Id)
	if err != nil {
		panic(err)
	}

	bookmarked := false
	for _, userBookmark := range userBookmarks {
		if userBookmark.DocumentId == d.Id {
			bookmarked = true
		}
	}

	referenced := false
	for _, userReference := range userReferences {
		if userReference.DocumentId == d.Id {
			referenced = true
		}
	}

	return &domain.Document{
		Id:         id,
		Title:      d.Title,
		File:       d.File,
		BookMarked: bookmarked,
		Referenced: referenced,
	}, nil
}
