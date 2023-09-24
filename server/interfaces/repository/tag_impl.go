package repository

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TagRepository struct {
	conn *gorm.DB
}

func NewTagRepository(conn *gorm.DB) *TagRepository {
	return &TagRepository{conn}
}

func (r *TagRepository) GetTags() ([]*domain.Tag, error) {
	var result []*domain.Tag

	if err := r.conn.Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}

func (r *TagRepository) CreateTag(id uuid.UUID, name string) error {
	tag := &domain.Tag{
		Id:   id,
		Name: name,
	}

	if err := r.conn.Create(tag).Error; err != nil {
		return err
	}

	return nil
}
