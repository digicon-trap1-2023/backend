package model

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/google/uuid"
)

type Document struct {
	Id          string `gorm:"type:char(36);not null;primaryKey"`
	UserId      string `gorm:"type:char(36);not null"`
	Title       string `gorm:"type:varchar(40)"`
	Description string `gorm:"type:varchar(200)"`
	File        string `gorm:"type:varchar(200)"`
}

func (Document) TableName() string {
	return "documents"
}

func (d *Document) ToDomain(userBookmarks []*BookMark, userReferences []*Reference, tags []*Tag) (*domain.Document, error) {
	id, err := uuid.Parse(d.Id)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(d.UserId)
	if err != nil {
		return nil, err
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

	tagsDomain := make([]*domain.Tag, len(tags))
	for i, tag := range tags {
		tagsDomain[i], err = tag.ToDomain()
		if err != nil {
			return nil, err
		}
	}

	return &domain.Document{
		Id:          id,
		UserId:      userId,
		Title:       d.Title,
		File:        d.File,
		Description: d.Description,
		Tags:        tagsDomain,
		BookMarked:  bookmarked,
		Referenced:  referenced,
	}, nil
}

func (d *Document) ToOtherDomain(userReferences []*Reference, tags []*Tag) (*domain.Document, error) {
	id, err := uuid.Parse(d.Id)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(d.UserId)
	if err != nil {
		return nil, err
	}

	referenceUsers := make([]string, 0)
	for _, userReference := range userReferences {
		if userReference.DocumentId == d.Id {
			referenceUsers = append(referenceUsers, userReference.UserId)
		}
	}

	tagsDomain := make([]*domain.Tag, len(tags))
	for i, tag := range tags {
		tagsDomain[i], err = tag.ToDomain()
		if err != nil {
			return nil, err
		}
	}

	return &domain.Document{
		Id:             id,
		UserId:         userId,
		Title:          d.Title,
		File:           d.File,
		Description:    d.Description,
		Tags:           tagsDomain,
		Referenced:     len(referenceUsers) > 0,
		ReferenceUsers: referenceUsers,
	}, nil
}
