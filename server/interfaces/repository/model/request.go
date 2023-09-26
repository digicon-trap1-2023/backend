package model

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/google/uuid"
)

type Request struct {
	Id          string `gorm:"type:char(36);not null;primaryKey"`
	Title       string `gorm:"type:varchar(36)"`
	Description string `gorm:"type:varchar(200)"`
	UserId      string `gorm:"type:char(36)"`
}

type TagRequest struct {
	RequestId string `gorm:"type:char(36);not null;primaryKey"`
	TagId     string `gorm:"type:char(36);not null;primaryKey"`
}

func (request *Request) ToDomain(tags []uuid.UUID) (*domain.Request, error) {
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	userId, err := uuid.Parse(request.UserId)
	if err != nil {
		return nil, err
	}

	return &domain.Request{
		Id:          id,
		Title:       request.Title,
		Description: request.Description,
		Tags:        tags,
		CreatedBy:   userId,
	}, nil
}
