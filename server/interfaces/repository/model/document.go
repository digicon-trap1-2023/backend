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
	FileWidth   int    `gorm:"type:int"`
	FileHeight  int    `gorm:"type:int"`
}

func (Document) TableName() string {
	return "documents"
}

func (d *Document) ToDomain(userBookmarks []*BookMark, userReferences []*Reference, tags []*Tag, userName string) (*domain.Document, error) {
	id, err := uuid.Parse(d.Id)
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
		UserName:    userName,
		Title:       d.Title,
		File:        d.File,
		FileHeight:  d.FileHeight,
		FileWidth:   d.FileWidth,
		Description: d.Description,
		Tags:        tagsDomain,
		BookMarked:  bookmarked,
		Referenced:  referenced,
	}, nil
}

func (d *Document) ToOtherDomain(userReferences []*Reference, userMap map[string]string, tags []*Tag) (*domain.Document, error) {
	id, err := uuid.Parse(d.Id)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(d.UserId)
	if err != nil {
		return nil, err
	}

	referenceUserIds := make([]string, 0)
	for _, userReference := range userReferences {
		if userReference.DocumentId == d.Id {
			referenceUserIds = append(referenceUserIds, userReference.UserId)
		}
	}

	referenceUsers := make([]string, 0)
	for _, referenceUserId := range referenceUserIds {
		referenceUsers = append(referenceUsers, userMap[referenceUserId])
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
		UserName:       userMap[userId.String()],
		Title:          d.Title,
		File:           d.File,
		FileHeight:     d.FileHeight,
		FileWidth:      d.FileWidth,
		Description:    d.Description,
		Tags:           tagsDomain,
		Referenced:     len(referenceUsers) > 0,
		ReferenceUsers: referenceUsers,
	}, nil
}
