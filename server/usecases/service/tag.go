package service

import (
	"github.com/digicon-trap1-2023/backend/domain"
	"github.com/digicon-trap1-2023/backend/interfaces/repository"
	"github.com/google/uuid"
)

type TagService struct {
	tagRepository *repository.TagRepository
}

func NewTagService(tagRepository *repository.TagRepository) *TagService {
	return &TagService{
		tagRepository: tagRepository,
	}
}

func (s *TagService) GetTags() ([]*domain.Tag, error) {
	return s.tagRepository.GetTags()
}

func (s *TagService) CreateTag(name string) (*domain.Tag, error) {
	tag := &domain.Tag{
		Id:   uuid.New(),
		Name: name,
	}

	if err := s.tagRepository.CreateTag(tag); err != nil {
		return nil, err
	}

	return tag, nil
}
