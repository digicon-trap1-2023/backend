package model

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/google/uuid"
)

type Tag struct {
	Id   string `gorm:"type:char(36);not null;primaryKey"`
	Name string `gorm:"type:varchar(40)"`
}

type TagDocument struct {
	TagId      string `gorm:"type:char(36);not null;primaryKey"`
	DocumentId string `gorm:"type:char(36);not null;primaryKey"`
}

func (tag *Tag) ToDomain() (*domain.Tag, error) {
	id, err := uuid.Parse(tag.Id)

	if err != nil {
		return nil, err
	}

	return &domain.Tag{
		Id:   id,
		Name: tag.Name,
	}, nil
}
